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
	var application models.Application

	// querySelectAllApplications = `SELECT app_acronym, app_description, app_Rnum, app_startDate, app_endDate FROM application`
	rows, err := middleware.SelectAllApplications()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var data []models.Application

	for rows.Next() {

		if err := rows.Scan(&application.AppAcronym, &application.Description, &application.Rnumber, &application.StartDate, &application.EndDate); err != nil {
			panic(err)
		}

		response := models.Application{
			AppAcronym:  application.AppAcronym,
			Description: application.Description,
			Rnumber:     application.Rnumber,
			StartDate:   application.StartDate,
			EndDate:     application.EndDate,
		}
		data = append(data, response)

	}
	c.JSON(200, data)
}

// route: /get-application
func GetApplication(c *gin.Context) {
	var application models.Application
	if err := c.BindQuery(&application); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}
	application.AppAcronym = c.Query("app_acronym")
	result := middleware.SelectSingleApplication(application.AppAcronym)

	switch err := result.Scan(&application.Description, &application.Rnumber, &application.PermitCreate, &application.PermitOpen, &application.PermitToDo, &application.PermitDoing, &application.PermitDone, &application.CreatedDate, &application.StartDate, &application.EndDate); err {

	// Create application
	// Date format yyyy-mm-dd
	case sql.ErrNoRows:
		middleware.ErrorHandler(c, 400, "Invalid app acronym")
		return
	}

	fmt.Println(application.StartDate)

	query := models.Application{
		AppAcronym:   application.AppAcronym,
		Description:  application.Description,
		Rnumber:      application.Rnumber,
		StartDate:    application.StartDate,
		EndDate:      application.EndDate,
		PermitCreate: application.PermitCreate,
		PermitOpen:   application.PermitOpen,
		PermitToDo:   application.PermitToDo,
		PermitDoing:  application.PermitDoing,
		PermitDone:   application.PermitDone,
		CreatedDate:  application.CreatedDate,
	}

	c.JSON(200, query)

}
