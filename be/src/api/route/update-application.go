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
	application.AppAcronym = c.Query("app_acronym")
	currentData := getSelectedApp(application.AppAcronym)

	if application.StartDate == "" {
		updateStartDate(application.AppAcronym, currentData["start_date"], c)
	} else if (application.StartDate != ""){
		updateStartDate(application.AppAcronym, application.StartDate, c)
	}

	if application.EndDate == "" {
		updateEndDate(application.AppAcronym, currentData["end_date"], c)
	} else if (application.EndDate != "") {
		updateEndDate(application.AppAcronym, application.EndDate, c)
	}

	if application.PermitCreate == "" {
		updatePermitCreate(application.AppAcronym, currentData["app_permitCreate"], c)
	} else if application.PermitCreate != "" {
		updatePermitCreate(application.AppAcronym, application.PermitCreate, c)
	}

	fmt.Println(application.PermitCreate)
	
	if application.PermitOpen == "" {
		updatePermitOpen(application.AppAcronym, currentData["app_permitOpen"], c)
	} else if application.PermitOpen != "" {
		updatePermitOpen(application.AppAcronym, application.PermitOpen, c)
	}

	if application.PermitToDo == "" {
		updatePermitToDo(application.AppAcronym, currentData["app_permitToDo"], c)
	} else if application.PermitToDo != "" {
		updatePermitToDo(application.AppAcronym, application.PermitToDo, c)
	}

	if application.PermitDoing == "" {
		updatePermitDoing(application.AppAcronym,currentData["app_permitDoing"], c)
	} else if application.PermitDoing != "" {
		updatePermitDoing(application.AppAcronym, application.PermitDoing, c)
	}

	if application.PermitDone == "" {
		updatePermitDone(application.AppAcronym, currentData["app_permitDone"], c)
	} else if application.PermitDone != "" {
		updatePermitDone(application.AppAcronym, application.PermitDone, c)
	}

	// result := middleware.SelectSingleApplication(application.AppAcronym)
	// fmt.Println(application.AppAcronym)
	// switch err := result.Scan(&application.Description, &application.Rnumber, &application.PermitCreate, &application.PermitOpen, &application.PermitToDo, &application.PermitDoing, &application.PermitDone, &application.CreatedDate, &application.StartDate, &application.EndDate); err {
	// //Application not found
	// case sql.ErrNoRows: 
	// 	middleware.ErrorHandler(c, 400, "Application does not exist")	
	// 	return
	
	// //Application exists
	// case nil:
	// 	_, err := middleware.UpdateApplication(application.StartDate, application.EndDate, application.PermitCreate, application.PermitOpen, application.PermitToDo, application.PermitDoing, application.PermitDone, application.AppAcronym)
	// 	fmt.Println(application.PermitCreate)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// c.JSON(200, gin.H{"code": 200, "message": "Application successfully updated!"})
}

func updateStartDate(AppAcronym string, StartDate string, c *gin.Context) {
	_, err := middleware.UpdateApplicationStartDate(StartDate, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("StartDate %s was successfully updated!", StartDate)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updateEndDate(AppAcronym string, EndDate string, c *gin.Context) {
	_, err := middleware.UpdateApplicationEndDate(EndDate, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("EndDate %s was successfully updated!", EndDate)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updatePermitCreate(AppAcronym string, PermitCreate string, c *gin.Context) {
	_, err := middleware.UpdateApplicationPermitCreate(PermitCreate, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("PermitCreate %s was successfully updated!", PermitCreate)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updatePermitOpen(AppAcronym string, PermitOpen string, c *gin.Context) {
	_, err := middleware.UpdateApplicationPermitOpen(PermitOpen, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("PermitOpen %s was successfully updated!", PermitOpen)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updatePermitToDo(AppAcronym string, PermitToDo string, c *gin.Context) {
	_, err := middleware.UpdateApplicationPermitToDo(PermitToDo, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("PermitToDo %s was successfully updated!", PermitToDo)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updatePermitDoing(AppAcronym string, PermitDoing string, c *gin.Context) {
	_, err := middleware.UpdateApplicationPermitDoing(PermitDoing, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("PermitDoing %s was successfully updated!", PermitDoing)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func updatePermitDone(AppAcronym string, PermitDone string, c *gin.Context) {
	_, err := middleware.UpdateApplicationPermitDone(PermitDone, AppAcronym)
	if err != nil {
		fmt.Println(err)
	}
	successMessage := fmt.Sprintf("PermitDone %s was successfully updated!", PermitDone)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
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
