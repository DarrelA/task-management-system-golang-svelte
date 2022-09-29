package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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
	var PermitCreate string
	result := middleware.SelectPermitCreate(task.TaskAppAcronym)
	err := result.Scan(&PermitCreate) 
	if (err != sql.ErrNoRows) {
		checkGroup := middleware.CheckGroup(c.GetString("username"), PermitCreate)
		if !checkGroup {
			middleware.ErrorHandler(c, 400, "Unauthorized actions")
			return
		} else {
			validateTaskName(task, c)
		}
	} else {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
	}
}

func validateTaskName(task models.Task, c *gin.Context) {
	var TaskName, TaskAppAcronym string
	if (!middleware.CheckLength(task.TaskName)) {
		result := middleware.SelectTaskName(task.TaskName, task.TaskAppAcronym)
		err := result.Scan(&TaskName, &TaskAppAcronym)
		if (err != sql.ErrNoRows) {
			error_message := fmt.Sprintf(`Task Name "%s" already exists for Application "%s"`, task.TaskName, task.TaskAppAcronym)
			middleware.ErrorHandler(c, 400, error_message)
		} else if (err == sql.ErrNoRows) {
			TaskID := generateTaskId(task, c)
			fmt.Println("from validatetaskname function:", TaskID)
			// c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Task was created!"})
		} else {
			checkError(err)
		}
	} else {
		middleware.ErrorHandler(c, 400, "Please enter a Task Name.")
	}
}

func generateTaskId(task models.Task, c *gin.Context) string {
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
	default:
		checkError(err)
	}
	return task.TaskID
}