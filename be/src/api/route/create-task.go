package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	validatePermitCreate(task, c)
}

func validatePermitCreate(task models.Task, c *gin.Context) {
	var PermitCreate sql.NullString
	result := middleware.SelectPermitCreate(task.TaskAppAcronym)
	err := result.Scan(&PermitCreate)
	if err != sql.ErrNoRows {
		checkGroup := middleware.CheckGroup(c.GetString("username"), PermitCreate.String)
		if !checkGroup {
			middleware.ErrorHandler(c, 400, "Unauthorized actions")
			return
		} else {
			task.TaskCreator = c.GetString("username")
			task.TaskOwner = c.GetString("username")
			validateTaskName(task, c)
		}
	} else {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
	}
}

func validateTaskName(task models.Task, c *gin.Context) {
	var TaskName, TaskAppAcronym sql.NullString
	if !middleware.CheckLength(task.TaskName) {
		task.TaskName = strings.TrimSpace(task.TaskName)
		result := middleware.SelectTaskName(task.TaskName, task.TaskAppAcronym)
		err := result.Scan(&TaskName, &TaskAppAcronym)
		if err != sql.ErrNoRows {
			error_message := fmt.Sprintf(`Task Name "%s" already exists for Application "%s"`, task.TaskName, task.TaskAppAcronym)
			middleware.ErrorHandler(c, 400, error_message)
		} else if err == sql.ErrNoRows {
			task.TaskID = generateTaskId(task, c)
			validateTaskPlan(task, c)
		} else {
			checkError(err)
		}
	} else {
		middleware.ErrorHandler(c, 400, "Please enter a Task Name.")
	}
}

func validateTaskPlan(task models.Task, c *gin.Context) {
	var PlanColor sql.NullString
	if !middleware.CheckLength(task.TaskPlan) {
		result := middleware.SelectPlanColor(task.TaskPlan, task.TaskAppAcronym)
		switch err := result.Scan(&PlanColor); {
		case err == sql.ErrNoRows:
			task.TaskColor = ""
			validateTaskNotes(task, c)
		case err != sql.ErrNoRows:
			task.TaskColor = PlanColor.String
			validateTaskNotes(task, c)
		default:
			checkError(err)
		}
	} else {
		validateTaskNotes(task, c)
	}
}

func validateTaskNotes(task models.Task, c *gin.Context) {
	task.TaskState = "Open"
	var TaskNotesDate, TaskNotesTime sql.NullString
	if !middleware.CheckLength(task.TaskNotes) {
		insertTaskTable(task)
		result := middleware.SelectCreatedTaskNotesTimestamp(task.TaskName, task.TaskAppAcronym)
		result.Scan(&TaskNotesDate, &TaskNotesTime)
		taskNotesAuditString := TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + task.TaskOwner + ", Task State: " + task.TaskState + "\n" + task.TaskNotes + " \n"
		_, err := middleware.UpdateTaskAuditNotes(taskNotesAuditString, task.TaskName, task.TaskAppAcronym)
		checkError(err)
		_, err = middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState, task.TaskAppAcronym)
		checkError(err)
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Task was successfully created!"})
	} else {
		insertTaskTable(task)
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Task was successfully created!"})
	}
}

func insertTaskTable(task models.Task) {
	var TaskPlan *string = nil
	if !middleware.CheckLength(task.TaskPlan) {
		_, err := middleware.InsertTask(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
	} else {
		_, err := middleware.InsertTaskWithoutPlan(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
	}
}

func generateTaskId(task models.Task, c *gin.Context) string {
	var TaskID sql.NullString
	var AppRNum sql.NullInt64
	result := middleware.SelectRNumber(task.TaskAppAcronym)
	result.Scan(&AppRNum)

	result = middleware.SelectTaskID(task.TaskAppAcronym)
	switch err := result.Scan(&TaskID); {
	case err == sql.ErrNoRows:
		// First task created in the application
		task.TaskID = task.TaskAppAcronym + "_" + strconv.Itoa(int(AppRNum.Int64))
	case err != sql.ErrNoRows:
		AppRNum.Int64 = AppRNum.Int64 + 1
		task.TaskID = task.TaskAppAcronym + "_" + strconv.Itoa(int(AppRNum.Int64))
		_, err := middleware.UpdateRNumber(int(AppRNum.Int64), task.TaskAppAcronym)
		checkError(err)
	default:
		checkError(err)
	}
	return task.TaskID
}
