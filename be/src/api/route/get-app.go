package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// route: /get-all-applications
// remove this comment once added
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
// remove this comment once added
