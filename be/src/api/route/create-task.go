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
	fmt.Println("validatePermitCreate")
	var PermitCreate string
	result := middleware.SelectPermitCreate(task.TaskAppAcronym)
	err := result.Scan(&PermitCreate) 
	if (err != sql.ErrNoRows) {
		checkGroup := middleware.CheckGroup(c.GetString("username"), PermitCreate)
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
	fmt.Println("validateTaskName")
	var TaskName, TaskAppAcronym string
	if (!middleware.CheckLength(task.TaskName)) {
		task.TaskName = strings.TrimSpace(task.TaskName)
		result := middleware.SelectTaskName(task.TaskName, task.TaskAppAcronym)
		err := result.Scan(&TaskName, TaskAppAcronym)
		if (err != sql.ErrNoRows) {
			error_message := fmt.Sprintf(`%s already exists in %s Application`, task.TaskName, task.TaskAppAcronym)
			middleware.ErrorHandler(c, 400, error_message)
		} else if (err == sql.ErrNoRows) {
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
	fmt.Println("validateTaskPlan")
	var PlanColor string
	if (!middleware.CheckLength(task.TaskPlan)) {
		result := middleware.SelectPlanColor(task.TaskPlan)
		switch err := result.Scan(&PlanColor); {
		case err == sql.ErrNoRows:
			task.TaskColor = ""
			validateTaskNotes(task, c)
		case err != sql.ErrNoRows:
			fmt.Println(PlanColor) 
			task.TaskColor = PlanColor
			validateTaskNotes(task, c)
		default:
			checkError(err)
		}
	} else {
		validateTaskNotes(task, c)
	}
}

func validateTaskNotes(task models.Task, c *gin.Context) {
	fmt.Println("validateTaskNotes")
	task.TaskState = "Open"
	var TaskNotesDate, TaskNotesTime string
	if (!middleware.CheckLength(task.TaskNotes)) {
		_, err := middleware.InsertTask(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
		result := middleware.SelectTaskNotesTimestamp(task.TaskName)
		result.Scan(&TaskNotesDate, &TaskNotesTime)
		taskNotesAuditString := TaskNotesDate + " " + TaskNotesTime + "\n" + "Task Owner: " + task.TaskOwner + ", Task State: " + task.TaskState  + "\n" + task.TaskNotes + " \n";
		fmt.Println("tasknotesAuditString: ", taskNotesAuditString)
		_, err = middleware.UpdateTaskAuditNotes(taskNotesAuditString, task.TaskName, task.TaskAppAcronym)
		checkError(err)
		_, err = middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState)
		checkError(err)
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Task was successfully created!"})
	} else {
		_, err := middleware.InsertTask(task.TaskAppAcronym, task.TaskID, task.TaskName, task.TaskDescription, task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskState, task.TaskCreator, task.TaskOwner)
		checkError(err)
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Task was successfully created!"})
	}
}

func generateTaskId(task models.Task, c *gin.Context) string {
	fmt.Println("generateTaskID")
	var TaskID string
	var AppRNum int
	result := middleware.SelectRNumber(task.TaskAppAcronym)
	result.Scan(&AppRNum)

	result = middleware.SelectTaskID(task.TaskAppAcronym)
	switch err := result.Scan(&TaskID); {
	case err == sql.ErrNoRows:
		// First task created in the application
		task.TaskID = task.TaskAppAcronym + "_" + strconv.Itoa(AppRNum)
	case err != sql.ErrNoRows:
		AppRNum = AppRNum + 1
		task.TaskID = task.TaskAppAcronym + "_" + strconv.Itoa(AppRNum)
		_, err := middleware.UpdateRNumber(AppRNum, task.TaskAppAcronym)
		checkError(err)
	default:
		checkError(err)
	}
	return task.TaskID
}