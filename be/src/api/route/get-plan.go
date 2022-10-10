package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// route: /get-all-plans?AppAcronym=durian
func GetAllPlans(c *gin.Context) {

	var pData []models.Plan
	var planName, planColor, planStartDate, planEndDate sql.NullString
	fmt.Println(c.GetString("username"))
	checkPM := middleware.CheckGroup(c.GetString("username"), "Project Manager")
	
	// AppAcronym URL params will be passed in here
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

	if len(pData) == 0 {
		fmt.Println("1")
		c.JSON(200, gin.H{"plans": []string{}, "checkPM": checkPM})
		return
	} else {
		c.JSON(200, gin.H{"plans": pData, "checkPM": checkPM})
	}

	
}
