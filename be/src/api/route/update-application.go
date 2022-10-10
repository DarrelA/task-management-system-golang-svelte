package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateApplication(c *gin.Context) {
	var application models.Application

	if err := c.BindQuery(&application); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := c.BindJSON(&application); err != nil {
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	fmt.Println(application)

	//Check User Group
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Project Lead")
	fmt.Println(c.GetString("username"))
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	application.AppAcronym = c.Query("AppAcronym")
	// currentData := getSelectedApp(application.AppAcronym)

	var (
		start *string = nil
		end   *string = nil
	)

	if application.StartDate == "" && application.EndDate == "" {
		_, err := middleware.UpdateApplicationNullDate(application.Description, start, end, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, application.AppAcronym)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Application successfully updated"})
		return
	}

	if application.StartDate == "" {
		fmt.Println("line 45")
		_, err := middleware.UpdateApplicationNullStartDate(application.Description, start, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, application.AppAcronym)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Application successfully updated"})
		return
	}

	if application.EndDate == "" {
		fmt.Println("line 55")
		_, err := middleware.UpdateApplicationNullEndDate(application.Description, application.StartDate, end, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, application.AppAcronym)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Application successfully updated"})
		return
	}

	_, err := middleware.UpdateApplication(application.Description, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, application.AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("line 68")
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Application successfully updated"})

}
