package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
) 

func UpdateApplication(c *gin.Context) {
	var application models.Application
	if err := c.BindQuery(&application); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}
	GetApplication(c)
	application.AppAcronym = c.Query("app_acronym")
	result := middleware.SelectSingleApplication(application.AppAcronym)
	fmt.Println(application.AppAcronym)
	switch err := result.Scan(&application.Description, &application.Rnumber, &application.PermitCreate, &application.PermitOpen, &application.PermitToDo, &application.PermitDoing, &application.PermitDone, &application.CreatedDate, &application.StartDate, &application.EndDate); err {
	//Application not found
	case sql.ErrNoRows: 
		middleware.ErrorHandler(c, 400, "Application does not exist")	
		return
	
	//Application exists
	case nil:
		_, err := middleware.UpdateApplication(application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone)
		if err != nil {
			fmt.Println(err)
			 
		}
	}

	c.JSON(200, gin.H{"code": 200, "message": "Application successfully updated!"})

}

                                 