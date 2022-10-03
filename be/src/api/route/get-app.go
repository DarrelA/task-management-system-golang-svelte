package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

// route: /get-all-applications
func GetAllApplications(c *gin.Context) {
	var (
		application models.Application
		acronym     sql.NullString
	)

	// querySelectAllApplications = `SELECT app_acronym, app_description, app_Rnum, app_startDate, app_endDate FROM application`
	rows, err := middleware.SelectAllApplications()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var data []models.Application

	for rows.Next() {

		if err := rows.Scan(&acronym, &application.Description, &application.Rnumber, &application.StartDate, &application.EndDate); err != nil {
			panic(err)
		}

		response := models.Application{
			AppAcronym:  acronym.String,
			Description: application.Description,
			Rnumber:     application.Rnumber,
			StartDate:   application.StartDate,
			EndDate:     application.EndDate,
		}
		data = append(data, response)

	}
	c.JSON(200, data)
	
	recipient := []string{"project_lead@tms.com", "project_lead2@tms.com"}
	middleware.SendMail(c, recipient, "team_member@tms.com", "alfred", "microservice 1")
}

// route: /get-application
func GetApplication(c *gin.Context) {
	var (
		application  models.Application
		permitCreate sql.NullString
		permitOpen   sql.NullString
		permitToDo   sql.NullString
		permitDoing  sql.NullString
		permitDone   sql.NullString
	)
	if err := c.BindQuery(&application); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}
	application.AppAcronym = c.Query("app_acronym")
	result := middleware.SelectSingleApplication(application.AppAcronym)

	switch err := result.Scan(&application.Description, &application.Rnumber, &permitCreate, &permitOpen, &permitToDo, &permitDoing, &permitDone, &application.CreatedDate, &application.StartDate, &application.EndDate); err {

	// Create application
	case sql.ErrNoRows:
		middleware.ErrorHandler(c, 400, "Invalid app acronym")
		return
	}

	query := models.Application{
		AppAcronym:   application.AppAcronym,
		Description:  application.Description,
		Rnumber:      application.Rnumber,
		StartDate:    application.StartDate,
		EndDate:      application.EndDate,
		PermitCreate: permitCreate.String,
		PermitOpen:   permitOpen.String,
		PermitToDo:   permitToDo.String,
		PermitDoing:  permitDoing.String,
		PermitDone:   permitDone.String,
		CreatedDate:  application.CreatedDate,
	}

	c.JSON(200, query)
}
