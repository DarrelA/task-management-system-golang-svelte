package middleware

import (
	"database/sql"
)

var (
	querySelectPermitCreate              = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`
	querySelectAppPermits                = `SELECT app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone FROM application WHERE app_acronym = ?`
	querySelectRNumber                   = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName                  = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskState                 = `SELECT task_state FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID                    = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectPlanColor                 = `SELECT plan_color FROM plan WHERE plan_mvp_name = ? AND plan_app_acronym = ?;`
	querySelectPlanColorByApp 			 = `SELECT plan_color FROM plan WHERE plan_color = ? AND plan_app_acronym = ?;`
	querySelectTaskNotes                 = `SELECT task_notes FROM task WHERE task_name = ? AND task_app_acronym = ?`
	querySelectAllApplications           = `SELECT app_acronym, app_description, app_Rnum, app_startDate, app_endDate FROM application`
	querySelectApplicationByAcronym      = `SELECT app_acronym FROM application WHERE app_acronym = ?`
	querySelectSingleApplication         = `SELECT app_description, app_Rnum, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate, CONVERT(app_startDate, DATE), CONVERT(app_endDate, DATE) FROM application WHERE app_acronym = ?`
	querySelectCreatedTaskNotesTimestamp = `SELECT DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate, TIME_FORMAT(task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_name = ? AND task_app_acronym = ?`
	querySelectTaskNotesTimestamp        = `SELECT DATE_FORMAT(last_updated, "%d/%m/%Y") as formattedDate, TIME_FORMAT(last_updated, "%H:%i:%s") as formattedTime, task_note, task_owner, task_state FROM task_notes WHERE task_name = ? AND task_app_acronym = ? ORDER BY last_updated DESC;`

	querySelectOneTask = `SELECT task_id, task_name, task_description, task_notes, task_plan,task_color, task_state, task_creator, task_owner,DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate,TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_name = ? AND task_app_acronym = ?;`

	querySelectAllPlans = `SELECT plan_mvp_name, plan_color, DATE_FORMAT(plan_startDate, "%d/%m/%Y") as formattedStartDate, DATE_FORMAT(plan_endDate, "%d/%m/%Y") as formattedEndDate FROM plan WHERE plan_app_acronym = ?`
	querySelectPlan     = `SELECT plan_mvp_name FROM plan WHERE plan_mvp_name = ? and plan_app_acronym = ?;`

	querySelectAllTasks         = `SELECT task_id, task_name, task_description, task_notes, task_plan, task_color, task_state, task_creator, task_owner, DATE_FORMAT(task_createDate, "%d/%m/%Y") as formattedDate, TIME_FORMAT(Task_createDate, "%H:%i:%s") as formattedTime FROM task WHERE task_app_acronym = ?;`
	querySelectEmailByUsername  = `SELECT email FROM accounts WHERE username = ?;`
	querySelectEmailByUserGroup = `SELECT username, email FROM accounts WHERE user_group LIKE CONCAT("%", ? , "%");`
)

var (
	queryUpdateRNumber        = `UPDATE application SET app_Rnum = ? WHERE app_acronym = ?;`
	queryUpdateTaskState      = `UPDATE task SET task_owner = ?, task_state = ? WHERE task_name = ? AND task_app_acronym = ?;`
	queryUpdateTaskAuditNotes = `UPDATE task SET task_notes = ? WHERE task_name = ? AND task_app_acronym = ?;`
	queryUpdateApplication    = `UPDATE application SET app_startDate = ?, app_endDate = ?, app_permitCreate = ?, app_permitOpen = ?, app_permitToDo = ?, app_permitDoing = ?, app_permitDone = ? WHERE app_acronym = ?;`
	queryUpdateTask           = `UPDATE task SET task_notes = ?, task_plan = ?, task_color = ?,  task_owner = ? WHERE task_name = ? AND task_app_acronym = ?`
)

var (
	queryInsertApplication = `INSERT INTO application (app_acronym, app_description, app_Rnum, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW()); `

	queryInsertPlan = `INSERT INTO plan (plan_mvp_name, plan_app_acronym, plan_color, plan_startDate, plan_endDate) VALUES (?,?,?,?,?);`

	queryInsertTask            = `INSERT INTO task (task_app_acronym, task_id, task_name, task_description, task_notes, task_plan, task_color, task_state, task_creator, task_owner, task_createDate) VALUES (?,?,?,?,?,?,?,?,?,?,now());`
	queryInsertCreateTaskNotes = `INSERT INTO task_notes (task_name, task_note, task_owner, task_state, task_app_acronym, last_updated) VALUES (?,?,?,?,?,now());`
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

func InsertCreateTaskNotes(TaskName string, TaskNote string, TaskOwner string, TaskState string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryInsertCreateTaskNotes, TaskName, TaskNote, TaskOwner, TaskState, TaskAppAcronym)
	return result, err
}

func InsertPlan(PlanMVPName string, PlanAppAcronym string, PlanColor string, PlanStartDate string, PlanEndDate string) (sql.Result, error) {
	result, err := db.Exec(queryInsertPlan, PlanMVPName, PlanAppAcronym, PlanColor, PlanStartDate, PlanEndDate)
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

func SelectTaskNotes(TaskName string, TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskNotes, TaskName, TaskAppAcronym)
	return result
}

func SelectTaskID(TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskID, TaskAppAcronym)
	return result
}

func SelectPlanColor(PlanName string, PlanAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectPlanColor, PlanName, PlanAppAcronym)
	return result
}

func SelectPlanColorByApp(PlanColor string, PlanAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectPlanColorByApp, PlanColor, PlanAppAcronym)
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

func SelectApplicationByAcronym(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectApplicationByAcronym, AppAcronym)
	return result
}

func SelectCreatedTaskNotesTimestamp(TaskName string, TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectCreatedTaskNotesTimestamp, TaskName, TaskAppAcronym)
	return result
}

func SelectTaskNotesTimestamp(TaskName string, TaskAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectTaskNotesTimestamp, TaskName, TaskAppAcronym)
	return result, err
}

func SelectAllPlans(PlanAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectAllPlans, PlanAppAcronym)
	return result, err
}

func SelectPlan(PlanMVPName string, PlanAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectPlan, PlanMVPName, PlanAppAcronym)
	return result
}

func SelectOneTask(TaskName string, TaskAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectOneTask, TaskName, TaskAppAcronym)
	return result, err
}

func SelectAllTasks(TaskAppAcronym string) (*sql.Rows, error) {
	result, err := db.Query(querySelectAllTasks, TaskAppAcronym)
	return result, err
}

func SelectEmailByUsername(username string) *sql.Row {
	result := db.QueryRow(querySelectEmailByUsername, username)
	return result
}

func SelectEmailByUserGroup(groupname string) (*sql.Rows, error) {
	result, err := db.Query(querySelectEmailByUserGroup, groupname)
	return result, err
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

func UpdateApplication(StartDate string, EndDate string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string, AppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateApplication, StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone, AppAcronym)
	return result, err
}

func UpdateTask(TaskNotes string, TaskPlan string, TaskPlanColor string, TaskOwner string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTask, TaskNotes, TaskPlan, TaskPlanColor, TaskOwner, TaskName, TaskAppAcronym)

	return result, err
}

func UpdateTaskWithoutPlan(TaskNotes string, TaskPlan *string, TaskPlanColor string, TaskOwner string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTask, TaskNotes, TaskPlan, TaskPlanColor, TaskOwner, TaskName, TaskAppAcronym)

	return result, err
}

// Insert Queries
func InsertApplication(AppAcronym string, Description string, Rnum int, StartDate string, EndDate string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string) (sql.Result, error) {
	result, err := db.Exec(queryInsertApplication, AppAcronym, Description, Rnum, StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone)
	return result, err
}

func InsertApplicationNullDate(AppAcronym string, Description string, Rnum int, StartDate *string, EndDate *string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string) (sql.Result, error) {
	result, err := db.Exec(queryInsertApplication, AppAcronym, Description, Rnum, StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone)
	return result, err
}

func UpdateTaskAuditNotes(TaskNotes string, TaskName string, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateTaskAuditNotes, TaskNotes, TaskName, TaskAppAcronym)

	return result, err
}
