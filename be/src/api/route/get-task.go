package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneTask(c *gin.Context) {
	// sql.NullString is a way to represent null string coming from SQL
	var TaskID, TaskDescription, TaskNotes, TaskPlan, TaskColor, TaskState, TaskCreator, TaskOwner, FormattedDate, FormattedTime sql.NullString

	TaskAppAcronym := c.Query("AppAcronym")
	TaskName := c.Query("TaskName")
	rows, err := middleware.SelectOneTask(TaskName, TaskAppAcronym)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Failed to /get-one-task")
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&TaskID, &TaskName, &TaskDescription, &TaskNotes, &TaskPlan, &TaskColor, &TaskState, &TaskCreator, &TaskOwner, &FormattedDate, &FormattedTime)
		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /get-one-task")
			return
		}

		data := models.Task{
			TaskAppAcronym:  TaskAppAcronym,
			TaskID:          TaskID.String,
			TaskName:        TaskName,
			TaskDescription: TaskDescription.String,
			TaskNotes:       TaskNotes.String,
			TaskPlan:        TaskPlan.String,
			TaskColor:       TaskColor.String,
			TaskState:       TaskState.String,
			TaskCreator:     TaskCreator.String,
			TaskOwner:       TaskOwner.String,
			FormattedDate:   FormattedDate.String,
			FormattedTime:   FormattedTime.String,
		}

		c.JSON(200, data)
	}

}

// http://localhost:4000/get-all-tasks/?AppAcronym=durian
func GetAllTasks(c *gin.Context) {
	var data []models.Task
	// sql.NullString is a way to represent null string coming from SQL
	var TaskID, TaskName, TaskDescription, TaskNotes, TaskPlan, TaskColor, TaskState, TaskCreator, TaskOwner, FormattedDate, FormattedTime sql.NullString
	// AppAcronym URL params will be passed in here
	TaskAppAcronym := c.Query("AppAcronym")
	rows, err := middleware.SelectAllTasks(TaskAppAcronym)
	if err != nil {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Failed to /get-all-tasks")
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&TaskID, &TaskName, &TaskDescription, &TaskNotes, &TaskPlan, &TaskColor, &TaskState, &TaskCreator, &TaskOwner, &FormattedDate, &FormattedTime)
		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Failed to scan in /get-all-tasks")
			return
		}

		response := models.Task{
			TaskAppAcronym:  TaskAppAcronym,
			TaskID:          TaskID.String,
			TaskName:        TaskName.String,
			TaskDescription: TaskDescription.String,
			TaskNotes:       TaskNotes.String,
			TaskPlan:        TaskPlan.String,
			TaskColor:       TaskColor.String,
			TaskState:       TaskState.String,
			TaskCreator:     TaskCreator.String,
			TaskOwner:       TaskOwner.String,
			FormattedDate:   FormattedDate.String,
			FormattedTime:   FormattedTime.String,
		}

		data = append(data, response)
	}

	c.JSON(200, data)
}
