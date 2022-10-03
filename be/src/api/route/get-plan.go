package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// route: /get-all-plans
func GetAllPlans(c *gin.Context) {
	// var plan models.Plan
	// if err := c.BindJSON(&plan); err != nil {
	//   return
	// }
	
	var pData []models.Plan
	var planName, planColor, planStartDate, planEndDate sql.NullString
	// AppAcronym URL params will be passed in here
	//var appAcro map[string][]string = c.Request.URL.Query()
	planAcronym := "test"
	rows, err := middleware.SelectAllPlans(planAcronym)
	
	if err != nil {
		log.Fatalln("select query error", err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&planName, &planColor, &planStartDate, &planEndDate)

		if err != nil {
			log.Fatalln("scan error", err)
		}

		resArray := models.Plan {
			PlanName: planName.String,
			PlanAcronym: planAcronym,
			PlanColor: planColor.String,
			StartDate: planStartDate.String,
			EndDate: planEndDate.String,
		}

		pData = append(pData, resArray)
	}

	c.JSON(200, pData)
	fmt.Println("get plans!")
}