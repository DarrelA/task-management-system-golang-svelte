package route

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"backend/api/middleware"
	"backend/api/models"

	"github.com/gin-gonic/gin"
)

func contains(s [5]string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func TaskStateTransition(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// validation for valid states
	validAppPermitStates := [5]string{"Open", "ToDo", "Doing", "Done", "Closed"}
	isValidState := contains(validAppPermitStates, task.TaskState)
	if !isValidState {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Invalid Task State")
		return
	}

	// validation for promoting and demoting state
	// after querying db TaskState
	taskState := middleware.SelectTaskState(task.TaskName, task.TaskAppAcronym)

	var TaskState sql.NullString
	err := taskState.Scan(&TaskState)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	if (TaskState.String == "Open" && task.TaskState != "Open" && task.TaskState != "ToDo") ||
		(TaskState.String == "ToDo" && task.TaskState != "ToDo" && task.TaskState != "Doing") ||
		(TaskState.String == "Doing" && task.TaskState != "Doing" && task.TaskState != "ToDo" && task.TaskState != "Done") ||
		(TaskState.String == "Done" && task.TaskState != "Done" && task.TaskState != "Doing" && task.TaskState != "Closed") ||
		(TaskState.String == "Closed" && task.TaskState != "Closed") {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Invalid State Transition")
		return
	}

	// validation for app permit rights
	// `UserGroups` is case sensitive
	appPermits := middleware.SelectAppPermits(task.TaskAppAcronym)

	var PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone sql.NullString
	err = appPermits.Scan(&PermitCreate, &PermitOpen, &PermitToDo, &PermitDoing, &PermitDone)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /task-state-transition")
		return
	}

	Username := c.GetString("username")
	userGroups := middleware.SelectUserFromUserGroupByUsername(Username)
	var UserGroups sql.NullString
	err = userGroups.Scan(&UserGroups)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /task-state-transition")
		return
	}

	switch checkAppPermit := TaskState.String; checkAppPermit {
	case "Open":
		isAuthorized := strings.Contains(UserGroups.String, PermitOpen.String)
		if !isAuthorized {
			middleware.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

	case "ToDo":
		isAuthorized := strings.Contains(UserGroups.String, PermitToDo.String)
		if !isAuthorized {
			middleware.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

	case "Doing":
		isAuthorized := strings.Contains(UserGroups.String, PermitDoing.String)
		if !isAuthorized {
			middleware.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

	case "Done":
		isAuthorized := strings.Contains(UserGroups.String, PermitDone.String)
		if !isAuthorized {
			middleware.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

	default:
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	_, err = middleware.UpdateTaskState(Username, task.TaskState, task.TaskName, task.TaskAppAcronym)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to update in /task-state-transition")
		return
	}

	if task.TaskState != TaskState.String {
		updateNotes := fmt.Sprintf("Task state has been updated from \"%s\" to \"%s\"", TaskState.String, task.TaskState)

		// insert task_note into TASK_NOTES TABLE when state is promoted or demoted
		_, err := middleware.InsertCreateTaskNotes(task.TaskName, updateNotes, Username, task.TaskState, task.TaskAppAcronym)
		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to insert notes in /task-state-transition")
			return
		}

		// insert task_note into TASK TABLE when state is promoted or demoted
		// logic from update-task.go
		rows, err := middleware.SelectTaskNotesTimestamp(task.TaskName, task.TaskAppAcronym)
		if err != nil {
			log.Fatal(err)
		}

		var TaskNotes, TaskNotesDate, TaskNotesTime, TaskOwner, TaskState sql.NullString
		var taskNotesAuditString string
		for rows.Next() {
			if err := rows.Scan(&TaskNotesDate, &TaskNotesTime, &TaskNotes, &TaskOwner, &TaskState); err != nil {
				log.Fatal(err)
			}

			taskNotesAuditString += TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + TaskOwner.String + ", Task State: " + TaskState.String + "\n" + TaskNotes.String + " \n\n"
		}

		_, err = middleware.UpdateTaskAuditNotes(taskNotesAuditString, task.TaskName, task.TaskAppAcronym)
		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to insert notes in /task-state-transition")
			return
		}
	}

	// query sender's email
	senderEmail := middleware.SelectEmailByUsername(Username)

	var SenderEmail sql.NullString

	err = senderEmail.Scan(&SenderEmail)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan senderEmail in /task-state-transition")
		return
	}

	// query list of Recipients to send emails
	// e.g. Get `emailList` of everyone who is in `Project Lead` usergroup
	// if `app_permitDone` under `application` table is `Project Lead`
	rows, err := middleware.SelectEmailByUserGroup(PermitDone.String)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to send email in /task-state-transition")
		return
	}
	defer rows.Close()

	// net/smtp in middleware.SendMail requires a slice of email(s)
	// calling the method x times the number of people in `app_permitDone` usergroup
	oneEmail := make([]string, 1)
	var RecipientUsername, RecipientEmail sql.NullString

	for rows.Next() {
		err = rows.Scan(&RecipientUsername, &RecipientEmail)
		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan recipients info in /task-state-transition")
			return
		}

		response := models.Recipients{
			RecipientUsername: RecipientUsername.String,
			RecipientEmail:    RecipientEmail.String,
		}

		// return null if no one is in `PermitDone.String` user_group
		oneEmail[0] = response.RecipientEmail

		// send email to ALL project leads if any from team member once task state is updated from `Doing` to `Done`
		if RecipientEmail.String != "" && TaskState.String == "Doing" && task.TaskState == "Done" {
			fmt.Println("middleware.SendMail called from task-state-transition.go")
			go middleware.SendMail(c, oneEmail, SenderEmail.String, Username, task.TaskName, response.RecipientUsername)
		}

	}

	c.JSON(200, gin.H{"message": "success"})
}

func GetUserAppPermits(c *gin.Context) {
	var task models.Task

	// validation for app permit rights
	// `UserGroups` is case sensitive
	task.TaskAppAcronym = c.Query("appacronym")
	appPermits := middleware.SelectAppPermits(task.TaskAppAcronym)

	var PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone sql.NullString
	err := appPermits.Scan(&PermitCreate, &PermitOpen, &PermitToDo, &PermitDoing, &PermitDone)
	if err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /task-state-transition")
		return
	}

	Username := c.GetString("username")
	userGroups := middleware.SelectUserFromUserGroupByUsername(Username)
	var UserGroups sql.NullString
	err = userGroups.Scan(&UserGroups)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /task-state-transition")
		return
	}

	type UserAppPermits struct {
		IsPermitCreate bool
		IsPermitOpen   bool
		IsPermitToDo   bool
		IsPermitDoing  bool
		IsPermitDone   bool
	}

	var userAppPermits UserAppPermits

	userAppPermits.IsPermitCreate = strings.Contains(UserGroups.String, PermitCreate.String)
	userAppPermits.IsPermitOpen = strings.Contains(UserGroups.String, PermitOpen.String)
	userAppPermits.IsPermitToDo = strings.Contains(UserGroups.String, PermitToDo.String)
	userAppPermits.IsPermitDoing = strings.Contains(UserGroups.String, PermitDoing.String)
	userAppPermits.IsPermitDone = strings.Contains(UserGroups.String, PermitDone.String)

	if PermitCreate.String == "" {
		userAppPermits.IsPermitCreate = false
	}
	if PermitOpen.String == "" {
		userAppPermits.IsPermitOpen = false
	}
	if PermitToDo.String == "" {
		userAppPermits.IsPermitToDo = false
	}
	if PermitDoing.String == "" {
		userAppPermits.IsPermitDoing = false
	}
	if PermitDone.String == "" {
		userAppPermits.IsPermitDone = false
	}

	c.JSON(200, gin.H{
		"IsPermitCreate": userAppPermits.IsPermitCreate,
		"IsPermitOpen":   userAppPermits.IsPermitOpen,
		"IsPermitToDo":   userAppPermits.IsPermitToDo,
		"IsPermitDoing":  userAppPermits.IsPermitDoing,
		"IsPermitDone":   userAppPermits.IsPermitDone,
	})
}
