package route

import (
	"database/sql"
	"fmt"
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

	var PermitOpen, PermitToDo, PermitDoing, PermitDone sql.NullString
	err = appPermits.Scan(&PermitOpen, &PermitToDo, &PermitDoing, &PermitDone)
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

	// insert note when state is promoted or demoted
	if task.TaskState != TaskState.String {
		updateNotes := fmt.Sprintf("Task state has been updated from \"%s\" to \"%s\"", TaskState.String, task.TaskState)
		_, err := middleware.InsertCreateTaskNotes(task.TaskName, updateNotes, Username, task.TaskState, task.TaskAppAcronym)
		if err != nil {
			fmt.Println(err)
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

	var emailList []string
	var RecipientUsername, RecipientEmail sql.NullString

	for rows.Next() {
		err = rows.Scan(&RecipientUsername, &RecipientEmail)
		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan recipients info in /task-state-transition")
			return
		}

		response := models.Recipients{
			RecipientUsername: RecipientUsername.String,
			RecipientEmail:    RecipientEmail.String,
		}

		// return null if no one is in `PermitDone.String` user_group
		emailList = append(emailList, response.RecipientEmail)
	}

	// send email to ALL project leads if any from team member once task state is updated from `Doing` to `Done`
	if RecipientEmail.String != "" && TaskState.String == "Doing" && task.TaskState == "Done" {
		fmt.Println("middleware.SendMail called from task-state-transition.go")
		// middleware.SendMail(c, emailList, SenderEmail.String, Username, task.TaskName)
	}

	// @TODO: discuss on what to return to FE
	// Below is for dev testing
	c.JSON(200, gin.H{
		"PermitOpen":     PermitOpen.String,
		"PermitToDo":     PermitToDo.String,
		"PermitDoing":    PermitDoing.String,
		"PermitDone":     PermitDone.String,
		"UserGroups":     UserGroups.String,
		"PermitDoneInfo": emailList,
	})
}
