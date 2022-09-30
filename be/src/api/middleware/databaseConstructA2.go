package middleware

import (
	"database/sql"
)

var (
	querySelectPermitCreate    = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`
	querySelectRNumber         = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName        = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID          = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectPlanColor       = `SELECT plan_color FROM plan WHERE plan_mvp_name = ?;`
	querySelectAllApplications = `SELECT app_acronym, app_description, app_Rnum, app_startDate, app_endDate FROM application`
	// querySelectSingleApplication = `SELECT app_description, app_Rnum, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate, DATE_FORMAT(app_startDate, "%Y-%m-%d") as app_startDate, DATE_FORMAT(app_endDate, "%Y-%m-%d") as app_endDate FROM application WHERE app_acronym = ?`
	querySelectSingleApplication = `SELECT app_description, app_Rnum, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate, CONVERT(app_startDate, DATE), CONVERT(app_endDate, DATE) FROM application WHERE app_acronym = ?`
)

var (
	queryUpdateRNumber = `UPDATE application SET app_Rnum = ? WHERE app_acronym = ?;`
)

var (
	queryInsertApplication = `INSERT INTO application (app_acronym, app_description, app_Rnum, app_startDate, app_endDate, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW()); `
)

// Select Queries
func SelectPermitCreate(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectPermitCreate, AppAcronym)
	return result
}

func SelectTaskName(TaskName string, TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskName, TaskName, TaskAppAcronym)
	return result
}

func SelectRNumber(AppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectRNumber, AppAcronym)
	return result
}

func SelectTaskID(TaskAppAcronym string) *sql.Row {
	result := db.QueryRow(querySelectTaskID, TaskAppAcronym)
	return result
}

func SelectPlanColor(TaskPlan string) *sql.Row {
	result := db.QueryRow(querySelectTaskID, TaskPlan)
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

// Update Queries
func UpdateRNumber(AppRNumber int, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateRNumber, AppRNumber, TaskAppAcronym)
	return result, err
}

// Insert Queries

func InsertApplication(AppAcronym string, Description string, Rnum int, StartDate string, EndDate string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string) (sql.Result, error) {
	result, err := db.Exec(queryInsertApplication, AppAcronym, Description, Rnum, StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone)
	return result, err
}
