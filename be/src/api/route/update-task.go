package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"net/http"

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
