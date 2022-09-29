package middleware

import "database/sql"

var (
	querySelectPermitCreate = `SELECT app_permitCreate FROM application WHERE app_acronym = ?`;
	querySelectRNumber = `SELECT app_Rnum FROM application WHERE app_acronym = ?;`
	querySelectTaskName = `SELECT task_name FROM task WHERE task_name = ? AND task_app_acronym = ?;`
	querySelectTaskID = `SELECT task_id FROM task WHERE task_app_acronym = ?;`
	querySelectPlanColor = `SELECT plan_color FROM plan WHERE plan_mvp_name = ?;`
)

var (
	queryUpdateRNumber = `UPDATE application SET app_Rnum = ? WHERE app_acronym = ?;`
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

// Update Queries
func UpdateRNumber(AppRNumber int, TaskAppAcronym string) (sql.Result, error) {
	result, err := db.Exec(queryUpdateRNumber, AppRNumber, TaskAppAcronym)
	return result, err
}