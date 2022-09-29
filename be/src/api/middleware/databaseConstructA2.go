package middleware

import "database/sql"

var (
	querySelectPermitCreate = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`;
	querySelectRNumber = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectSingleApplication = `SELECT app_description, app_Rnum, app_permitCreate, app_permitOpen, app_permitToDo, app_permitDoing, app_permitDone, app_createdDate, DATE_FORMAT(app_startDate, "%Y-%m-%d") as startDate, DATE_FORMAT(app_endDate, "%Y-%m-%d") as endDate FROM application WHERE app_acronym = ?;`
)

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

func SelectSingleApplication (AppAcronym string) * sql.Row {
	result := db.QueryRow(querySelectSingleApplication, AppAcronym)
	return result
}