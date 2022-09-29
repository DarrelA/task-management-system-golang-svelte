package models

import "database/sql"

type Application struct {
	AppAcronym   string `json: "app_acronym"`
	Description  string `json:"app_description"`
	Rnumber      string `json:"app_Rnum"`
	StartDate    string `json: "start_date"`
	EndDate      string `json: "end_date"`
	PermitCreate string `json: "app_permitCreate"`
	PermitOpen   string `json: "app_permitOpen"`
	PermitToDo   string `json: "app_permitTodo"`
	PermitDoing  string `json: "app_permitDoing"`
	PermitDone   string `json: "app_permitDone"`
	CreatedDate  string `json: "app_created"`
}

type Task struct {
	TaskAppAcronym      string `json:"task_app_acronym"`
	TaskID          string `json:"task_id"`
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	TaskNotes       string `json:"task_notes"`
	TaskPlan        string `json:"task_plan"`
	TaskColor       string `json:"task_color"`
	TaskState       string `json:"task_state"`
	TaskCreator     string `json:"task_creator"`
	TaskOwner       string `json:"task_owner"`
	CreatedDate     string `json:"task_created"`
	FormattedDate string `json:"formatted_date"`
	FormattedTime string `json:"formatted_time"`
}

type Plan struct {
	PlanName    string `json: "plan_name`
	PlanAcronym string `json: "plan_acronym"`
	PlanColor   string `json: "plan_color`
	StartDate   string `json: "plan_start"`
	EndDate     string `json: "plan_end"`
}
