package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// route: /get-all-plans?AppAcronym=durian
func GetAllPlans(c *gin.Context) {
	// var plan models.Plan
	// if err := c.BindJSON(&plan); err != nil {
	//   return
	// }

	var pData []models.Plan
	var planName, planColor, planStartDate, planEndDate sql.NullString
	// AppAcronym URL params will be passed in here
	//var appAcro map[string][]string = c.Request.URL.Query()
	planAppAcronym := c.Query("AppAcronym")
	rows, err := middleware.SelectAllPlans(planAppAcronym)

	if err != nil {
		middleware.ErrorHandler(c, http.StatusInternalServerError, "Error in querySelectAllPlans")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&planName, &planColor, &planStartDate, &planEndDate)

		if err != nil {
			middleware.ErrorHandler(c, http.StatusInternalServerError, "Error in scanning error in /get-all-plans")
			log.Fatalln("scan error", err)
		}

		resArray := models.Plan{
			PlanName:    planName.String,
			PlanAcronym: planAppAcronym,
			PlanColor:   planColor.String,
			StartDate:   planStartDate.String,
			EndDate:     planEndDate.String,
		}

		pData = append(pData, resArray)
	}

	c.JSON(200, pData)
}
