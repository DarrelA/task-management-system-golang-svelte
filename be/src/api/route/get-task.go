package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// WARNING: This is POST request because have to take AppAcronym
// from the params in the URL for tasks.
// http://localhost:4000/get-all-tasks/?AppAcronym=durian
func GetAllTasks(c *gin.Context) {
	var task models.Task
	var data []models.Task
	var formattedDate, formattedTime string
	// AppAcronym URL params will be passed in here
	var AppAcronym map[string][]string = c.Request.URL.Query()
	TaskAppAcronym := strings.Join(AppAcronym["AppAcronym"], "")
	rows, err := middleware.SelectAllTasks(TaskAppAcronym)
	checkGetError("Failed to /get-all-tasks: ", err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&task.TaskID, &task.TaskName, &task.TaskDescription, &task.TaskNotes, &task.TaskPlan, &task.TaskColor, &task.TaskState, &task.TaskCreator, &task.TaskOwner, &formattedDate, &formattedTime)
		checkGetError("Failed to scan in /get-all-tasks", err)

		response := models.Task {
			TaskAppAcronym: TaskAppAcronym,
			TaskID: task.TaskID,
			TaskName: task.TaskName,
			TaskDescription: task.TaskDescription,
			TaskNotes: task.TaskNotes,
			TaskPlan: task.TaskPlan,
			TaskColor: task.TaskColor,
			TaskState: task.TaskState,
			TaskCreator: task.TaskCreator,
			TaskOwner: task.TaskOwner,
			FormattedDate: formattedDate,
			FormattedTime: formattedTime,
		}

		data = append(data, response)
	}

	c.JSON(200, data)
	
	err = rows.Err()
	checkGetError("Some other error occurred", err)
}

// route: /get-task
// remove this comment once added

func checkGetError(message string, err error) {
	if err != nil {
		log.Fatalln(message, err)
	}
}