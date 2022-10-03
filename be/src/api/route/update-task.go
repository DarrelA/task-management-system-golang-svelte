package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TO BE DONE BY BEATRICE
// Update Task (with/without plan)
// Validation:
// - check task plan - if there is selected plan (check plan color - update color based on selected plan [if plan is empty, update plan color to empty string])
// - check task notes if there is task notes (insert in tasknotes table and update tasknotes in task table)/there is no task notes (dont need insert in tasknotes table and update tasknotes in task table)
func UpdateTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

}

// check if plan is empty
func CheckTaskPlan(task models.Task, c *gin.Context) {

}

// check if task notes is empty
func CheckTaskNotes(task models.Task, c *gin.Context) {

}

// insert task notes
func InsertTaskNotes(task models.Task, c *gin.Context) {
	// do all the checking
	// insert task notes
}

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
	// UserGroups is case sensitive
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

	// @TODO: update notes to include state change

	// @TODO: send email to ALL project leads from team member once task state is updated to done

	// @TODO: discuss on what to return to FE
	// Below is for dev testing
	c.JSON(200, gin.H{
		"PermitOpen":  PermitOpen.String,
		"PermitToDo":  PermitToDo.String,
		"PermitDoing": PermitDoing.String,
		"PermitDone":  PermitDone.String,
		"UserGroups":  UserGroups.String,
	})

}
