package middleware

import "database/sql"

var (
	queryInsertTask = `INSERT INTO task (task_app_acronym, task_id, task_name, task_description, task_notes, task_plan, task_color, task_state, task_creator, task_owner, task_createDate) VALUES (?,?,?,?,?,?,?,?,?,?,now());`
	queryInsertCreateTaskNotes = `INSERT INTO task_notes (task_name, task_note, task_owner, task_state, last_updated) VALUES (?,?,?,?,now());`
)

var (
	querySelectPermitCreate = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`;
	querySelectRNumber = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectPlanColor = `SELECT plan_color FROM plan WHERE plan_mvp_name = ?;`
	querySelectTaskNotesTimestamp = `SELECT DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate, TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_name = ?;`
)

var (
	queryUpdateRNumber = `UPDATE application SET app_Rnum = ? WHERE app_acronym = ?;`
	queryUpdateTaskAuditNotes = `UPDATE task SET task_notes = ? WHERE task_name = ? AND task_app_acronym = ?;`
)

// Insert Queries
func InsertTask(TaskAppAcronym string, TaskID string, TaskName string, TaskDescription string, TaskNotes string, TaskPlan string, TaskColor string, TaskState string, TaskCreator string, TaskOwner string) (sql.Result, error) {
	result, err := db.Exec(queryInsertTask, TaskAppAcronym, TaskID, TaskName, TaskDescription, TaskNotes, TaskPlan, TaskColor, TaskState, TaskCreator, TaskOwner)
	return result, err
}

func InsertCreateTaskNotes(TaskName string, TaskNote string, TaskOwner string, TaskState string) (sql.Result, error) {
	result, err := db.Exec(queryInsertCreateTaskNotes, TaskName, TaskNote, TaskOwner, TaskState)
	return result, err
}

// Select Queries
func SelectPermitCreate(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectPermitCreate, AppAcronym)
	return result
}

func SelectRNumber(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectRNumber, AppAcronym)
	return result
}

func SelectTaskName(TaskName string, TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskName, TaskName, TaskAppAcronym)
	return result
}

func SelectTaskID(TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskID, TaskAppAcronym)
	return result
}

func SelectPlanColor(TaskPlan string) *sql.Row {
	result := db.QueryRow(querySelectPlanColor, TaskPlan)
	return result
}

func SelectTaskNotesTimestamp(TaskName string) *sql.Row {
	result := db.QueryRow(querySelectTaskNotesTimestamp, TaskName)
	return result
}

// Update Queries
func UpdateRNumber(AppRNumber int, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateRNumber, AppRNumber, TaskAppAcronym)
	return result, err
}

func UpdateTaskAuditNotes(TaskNotes string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTaskAuditNotes, TaskNotes, TaskName, TaskAppAcronym)
	return result, err
}