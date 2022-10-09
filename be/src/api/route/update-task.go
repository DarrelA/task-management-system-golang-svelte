package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"
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
	// params
	task.TaskAppAcronym = c.Query("AppAcronym")
	task.TaskName = c.Query("TaskName")

	// call BindJSON to bind the received JSON to task
	if err := c.BindJSON(&task); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// check if taskname exists
	result := middleware.SelectTaskName(task.TaskName, task.TaskAppAcronym)

	fmt.Println(task.TaskAppAcronym)
	fmt.Println(task.TaskName)

	switch err := result.Scan(&task.TaskName); err {

	// task name does not exist
	case sql.ErrNoRows:
		middleware.ErrorHandler(c, 400, "Task name does not exist")
		return

	// task name exists
	case nil:
		// task state check
		task.TaskState = checkTaskState(task, c)

		// permit check
		task.TaskOwner = checkPermit(task, c)
		fmt.Println(task.TaskOwner)

		// update task with/without plan
		if (task.TaskOwner != "") {
			fmt.Println("line 56")
			
			// plan color check
			task.TaskColor = checkTaskPlanColor(task, c)

			// task notes check
			task.TaskNotes = checkTaskNotes(task, c)
			updateTaskTable(task, c)
			return
		}
	}
}

// check task state
// 1. no task state (return empty string)
// 2. has task state (return task state string)
func checkTaskState(task models.Task, c *gin.Context) string {
	var TaskState sql.NullString

	result := middleware.SelectTaskState(task.TaskName, task.TaskAppAcronym)

	switch err := result.Scan(&TaskState); err {
	case sql.ErrNoRows:
		task.TaskState = ""

	case nil:
		task.TaskState = TaskState.String
	}

	return task.TaskState
}



func checkPermit(task models.Task, c *gin.Context) string {
	var PermitCreate sql.NullString
	var PermitOpen sql.NullString
	var PermitToDo sql.NullString
	var PermitDoing sql.NullString
	var PermitDone sql.NullString
	var TaskOwner string

	// select query from application (5 permits)
	result := middleware.SelectAppPermits(task.TaskAppAcronym)

	switch err := result.Scan(&PermitCreate, &PermitOpen, &PermitToDo, &PermitDoing, &PermitDone); err {
	
	// no app permits
	case sql.ErrNoRows:
		middleware.ErrorHandler(c, 400, "Application does not exist")

	// app permits
	case nil:
		// for each task state, check user group if authorized to update task
		switch {
		case task.TaskState == "Open":
			checkGroup := middleware.CheckGroup(c.GetString("username"), PermitOpen.String)
			fmt.Println("username here", c.GetString("username"))
			fmt.Println(checkGroup)
			if !checkGroup {
				fmt.Println("Not Here")
				middleware.ErrorHandler(c, 400, "Unauthorized actions")
				TaskOwner = ""
			} else {
				fmt.Println("Here")
				TaskOwner = c.GetString("username")
			}
			// return TaskOwner

		case task.TaskState == "ToDo":
			checkGroup := middleware.CheckGroup(c.GetString("username"), PermitToDo.String)
			if !checkGroup {
				middleware.ErrorHandler(c, 400, "Unauthorized actions")
				TaskOwner = ""
			} else {
				TaskOwner = c.GetString("username")
			}
			// return TaskOwner

		case task.TaskState == "Doing":
			checkGroup := middleware.CheckGroup(c.GetString("username"), PermitDoing.String)
			if !checkGroup {
				middleware.ErrorHandler(c, 400, "Unauthorized actions")
				TaskOwner = ""
			} else {
				TaskOwner = c.GetString("username")
			}
			// return TaskOwner

		case task.TaskState == "Done":
			checkGroup := middleware.CheckGroup(c.GetString("username"), PermitDone.String)
			if !checkGroup {
				middleware.ErrorHandler(c, 400, "Unauthorized actions")
				TaskOwner = ""
			} else {
				TaskOwner = c.GetString("username")
			}
			// return TaskOwner

		default:
			fmt.Println("Task state not available")
			TaskOwner = ""
		}
	}

	return TaskOwner
}

// check if there is plan color
// 1. no plan color (return empty string)
// 2. has plan color (return plan color)
func checkTaskPlanColor(task models.Task, c *gin.Context) string {
	var PlanColor sql.NullString

	result := middleware.SelectPlanColor(task.TaskPlan, task.TaskAppAcronym)

	switch err := result.Scan(&PlanColor); err {
	case sql.ErrNoRows:
		task.TaskColor = ""

	case nil:
		task.TaskColor = PlanColor.String
	}

	return task.TaskColor
}

// check if there is task notes
// 1. no new task notes (get existing task notes from task table and return task notes)
// 2. has new task notes (insert task notes into task_notes table and return formatted task notes)
func checkTaskNotes(task models.Task, c *gin.Context) string {
	var TaskNotes, TaskNotesDate, TaskNotesTime, TaskOwner, TaskState sql.NullString
	var taskNotesAuditString string

	if !middleware.CheckLength(task.TaskNotes) {

		// insert task notes, task owner and task state into task notes table
		_, err := middleware.InsertCreateTaskNotes(task.TaskName, task.TaskNotes, task.TaskOwner, task.TaskState, task.TaskAppAcronym)

		if err != nil {
			panic(err)
		}

		// format new task notes
		// concat with existing task notes
		rows, err := middleware.SelectTaskNotesTimestamp(task.TaskName, task.TaskAppAcronym)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			if err := rows.Scan(&TaskNotesDate, &TaskNotesTime, &TaskNotes, &TaskOwner, &TaskState); err != nil {
				log.Fatal(err)
			}

			taskNotesAuditString += TaskNotesDate.String + " " + TaskNotesTime.String + "\n" + "Task Owner: " + TaskOwner.String + ", Task State: " + TaskState.String + "\n" + TaskNotes.String + " \n\n"
		}

		task.TaskNotes = taskNotesAuditString
		
	} else {
		// get existing task notes

		result := middleware.SelectTaskNotes(task.TaskName, task.TaskAppAcronym)

		switch err := result.Scan(&TaskNotes); err {

		case sql.ErrNoRows:
			middleware.ErrorHandler(c, 400, "Task does not exist")

		case nil:
			task.TaskNotes = TaskNotes.String
		}
	}

	return task.TaskNotes
}

// check if there is a plan
// 1. yes: insert with plan (plan name)
// 2. no: insert without plan (null)
func updateTaskTable(task models.Task, c *gin.Context) {
	var TaskPlan *string = nil

	if !middleware.CheckLength(task.TaskPlan) {
		_, err := middleware.UpdateTask(task.TaskNotes, task.TaskPlan, task.TaskColor, task.TaskOwner, task.TaskName, task.TaskAppAcronym)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("line 245")
		fmt.Println("line 246", task.TaskNotes)
		_, err := middleware.UpdateTaskWithoutPlan(task.TaskNotes, TaskPlan, task.TaskColor, task.TaskOwner, task.TaskName, task.TaskAppAcronym)
		if err != nil {
			panic(err)
		}
	}

	c.JSON(200, gin.H{"code": 200, "message": "Task updated successfully"})
}
