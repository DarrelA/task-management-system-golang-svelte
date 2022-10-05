package route

import (
	"database/sql"
	"fmt"

	"backend/api/middleware"
	"backend/api/models"

	"github.com/gin-gonic/gin"
)

func CreateApplication(c *gin.Context) {
	var (
		application models.Application
		acronym     sql.NullString
		start       *string = nil
		end         *string = nil
	)

	if err := c.BindJSON(&application); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}

	// Check group : Project Lead
	// replace GetString with "admin" for POSTMAN testing
	// checkGroup := middleware.CheckGroup(c.GetString("username"), "Project Lead")
	checkGroup := middleware.CheckGroup("admin", "Project Lead")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	if len(application.AppAcronym) == 0 {
		middleware.ErrorHandler(c, 400, "Invalid app acronym")
		return
	}

	if len(application.Description) == 0 {
		middleware.ErrorHandler(c, 400, "Input app description")
		return
	}

	if application.Rnumber <= 0 {
		middleware.ErrorHandler(c, 400, "Invalid app running number")
		return
	}

	// Query if acronym exist
	result := middleware.SelectApplicationByAcronym(application.AppAcronym)
	switch err := result.Scan(&acronym); err {
	// Create application
	case sql.ErrNoRows:
		if application.StartDate == "" {
			_, err := middleware.InsertApplicationNullDate(application.AppAcronym, application.Description, application.Rnumber, start, end, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone)
			if err != nil {
				middleware.ErrorHandler(c, 400, "Invalid request")
				fmt.Println(err)
				return
			}
		} else {
			_, err := middleware.InsertApplication(application.AppAcronym, application.Description, application.Rnumber, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone)
			if err != nil {
				middleware.ErrorHandler(c, 400, "Invalid request")
				fmt.Println(err)
				return
			}
		}

	case nil:
		middleware.ErrorHandler(c, 400, "Invalid app acronym")
		return
	}

	c.JSON(200, gin.H{"code": 200, "message": "Successfully added new application"})
}
