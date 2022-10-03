package middleware

import (
	"database/sql"
)

var (
	querySelectPermitCreate       = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`
	querySelectAppPermits         = `SELECT app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone FROM application WHERE app_acronym = ?`
	querySelectRNumber            = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName           = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskState          = `SELECT task_state FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID             = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectPlanColor          = `SELECT plan_color FROM plan WHERE plan_mvp_name = ?;`
	querySelectAllApplications    = `SELECT app_acronym, app_description, app_Rnum, app_startDate, app_endDate FROM application`
	querySelectSingleApplication  = `SELECT app_description, app_Rnum, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate, CONVERT(app_startDate, DATE), CONVERT(app_endDate, DATE) FROM application WHERE app_acronym = ?`
	querySelectTaskNotesTimestamp = `SELECT DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate, TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_name = ?;`

	querySelectAllPlans = `SELECT plan_mvp_name, plan_color, DATE_FORMAT(plan_startDate, "%d/%m/%Y") as formattedStartDate, DATE_FORMAT(plan_endDate, "%d/%m/%Y") as formattedEndDate FROM plan WHERE plan_app_acronym = ?`

	querySelectOneTask = `SELECT task_id, task_name, task_description, task_notes, task_plan,task_color, task_state, task_creator, task_owner,DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate,TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_name = ? AND task_app_acronym = ?;`

	querySelectAllTasks         = `SELECT task_id, task_name, task_description, task_notes, task_plan, task_color, task_state, task_creator, task_owner, DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate, TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_app_acronym = ?;`
	querySelectEmailByUserGroup = `SELECT accounts.email, usergroup.user_group FROM accounts, usergroup WHERE accounts.username = usergroup.username AND usergroup.user_group = ?`
)

var (
	queryUpdateRNumber        = `UPDATE application SET app_Rnum = ? WHERE app_acronym = ?;`
	queryUpdateTaskState      = `UPDATE task SET task_owner = ?, task_state = ? WHERE task_name = ? AND task_app_acronym = ?;`
	queryUpdateTaskAuditNotes = `UPDATE task SET task_notes = ? WHERE task_name = ? AND task_app_acronym = ?;`
)

var (
	queryInsertApplication = `INSERT INTO application (app_acronym, app_description, app_Rnum, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW()); `

	queryInsertTask            = `INSERT INTO task (task_app_acronym, task_id, task_name, task_description, task_notes, task_plan, task_color, task_state, task_creator, task_owner, task_createDate) VALUES (?,?,?,?,?,?,?,?,?,?,now());`
	queryInsertCreateTaskNotes = `INSERT INTO task_notes (task_name, task_note, task_owner, task_state, last_updated) VALUES (?,?,?,?,now());`
)

// Insert Queries
func InsertTask(TaskAppAcronym string, TaskID string, TaskName string, TaskDescription string, TaskNotes string, TaskPlan string, TaskColor string, TaskState string, TaskCreator string, TaskOwner string) (sql.Result, error) {
	result, err := db.Exec(queryInsertTask, TaskAppAcronym, TaskID, TaskName, TaskDescription, TaskNotes, TaskPlan, TaskColor, TaskState, TaskCreator, TaskOwner)
	return result, err
}

func InsertTaskWithoutPlan(TaskAppAcronym string, TaskID string, TaskName string, TaskDescription string, TaskNotes string, TaskPlan *string, TaskColor string, TaskState string, TaskCreator string, TaskOwner string) (sql.Result, error) {
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

func SelectAppPermits(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectAppPermits, AppAcronym)
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

func SelectTaskState(TaskName string, TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskState, TaskName, TaskAppAcronym)
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

func SelectAllApplications() (*sql.Rows, error) {
	result, err := db.Query(querySelectAllApplications)
	return result, err
}

func SelectSingleApplication(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectSingleApplication, AppAcronym)
	return result
}

func SelectTaskNotesTimestamp(TaskName string) *sql.Row {
	result := db.QueryRow(querySelectTaskNotesTimestamp, TaskName)
	return result
}

func SelectAllPlans(PlanAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectAllPlans, PlanAppAcronym)
	return result, err
}

func SelectOneTask(TaskName string, TaskAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectOneTask, TaskName, TaskAppAcronym)
	return result, err
}

func SelectAllTasks(TaskAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectAllTasks, TaskAppAcronym)
	return result, err
}

func SelectEmailByUserGroup() {

}

// Update Queries
func UpdateRNumber(AppRNumber int, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateRNumber, AppRNumber, TaskAppAcronym)
	return result, err
}

func UpdateTaskState(Username string, TaskState string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTaskState, Username, TaskState, TaskName, TaskAppAcronym)
	return result, err
}

// Insert Queries

func InsertApplication(AppAcronym string, Description string, Rnum int, StartDate string, EndDate string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string) (sql.Result, error) {
	result, err := db.Exec(queryInsertApplication, AppAcronym, Description, Rnum, StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone)
	return result, err
}

func UpdateTaskAuditNotes(TaskNotes string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTaskAuditNotes, TaskNotes, TaskName, TaskAppAcronym)

	return result, err
}
