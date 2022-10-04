package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
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

	//Check User Group
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Project Lead")
	fmt.Println(c.GetString("username"))
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	application.AppAcronym = c.Query("app_acronym")
	currentData := getSelectedApp(application.AppAcronym)

	/////////// Start Date ///////////
	if application.StartDate == "" {
		updateAppToDB(application.AppAcronym, currentData["start_date"], application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	} else if (application.StartDate != ""){
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// End Date ///////////
	if application.EndDate == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, currentData["end_date"], application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	} else if (application.EndDate != "") {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// Permit Create ///////////
	if application.PermitCreate == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, currentData["app_permitCreate"], application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	} else if application.PermitCreate != "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// Permit Open ///////////
	if application.PermitOpen == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, currentData["app_permitOpen"], application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	} else if application.PermitOpen != "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// Permit ToDo ///////////
	if application.PermitToDo == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, currentData["app_permitToDo"], application.PermitDoing, application.PermitDone, c)
	} else if application.PermitToDo != "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// Permit Doing ///////////
	if application.PermitDoing == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, currentData["app_permitDoing"], application.PermitDone, c)
	} else if application.PermitDoing != "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}

	/////////// Permit Done ///////////
	if application.PermitDone == "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, currentData["app_permitDone"], c)
	} else if application.PermitDone != "" {
		updateAppToDB(application.AppAcronym, application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, c)
	}
}

func updateAppToDB(AppAcronym string, StartDate string, EndDate string, PermitCreate string, PermitOpen string, PermitToDo string, PermitDoing string, PermitDone string, c *gin.Context) {
	_, err := middleware.UpdateApplication(StartDate, EndDate, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": "Application successfully updated"})
}

func getSelectedApp (AppAcronym string) map[string]string {
	var Description, PermitCreate, PermitOpen, PermitToDo, PermitDoing, PermitDone, StartDate, EndDate, CreatedDate string
	var Rnumber int
	result := middleware.SelectSingleApplication(AppAcronym)

	currentAppData := make(map[string]string)
	err := result.Scan(&Description, &Rnumber, &PermitCreate, &PermitOpen, &PermitToDo, &PermitDoing, &PermitDone, &CreatedDate, &StartDate, &EndDate)
	if err != sql.ErrNoRows {
		currentAppData["start_date"] = StartDate
		currentAppData["end_date"] = EndDate
		currentAppData["app_permitCreate"] = PermitCreate
		currentAppData["app_permitOpen"] = PermitOpen
		currentAppData["app_permitToDo"] = PermitToDo
		currentAppData["app_permitDoing"] = PermitDoing
		currentAppData["app_permitDone"] = PermitDone	
	}		
	return currentAppData
}
