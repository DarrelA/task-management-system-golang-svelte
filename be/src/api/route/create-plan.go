package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.BindJSON(&plan); err != nil {
		return
	}
	// static groupname or use permit open?
	checkPermit := middleware.CheckGroup(c.GetString("username"), "Project Manager")
	if checkPermit {
		//proceed with create plan
		if plan.PlanAcronym == "" || plan.PlanName == "" {
			middleware.ErrorHandler(c, 400, "Please enter a Plan Name!")
			return
		}

		if plan.StartDate == "" {
			middleware.ErrorHandler(c, 400, "Please enter a Start Date!")
			return
		}

		if plan.EndDate == "" {
			middleware.ErrorHandler(c, 400, "Please enter an End Date!")
			return
		}

		plan.PlanName = strings.TrimSpace(plan.PlanName)
		createPlan(plan, c)
	} else {
		middleware.ErrorHandler(c, 400, "Unauthorized to create plan")
		return
	}
}

func createPlan(plan models.Plan, c *gin.Context) {
	var PlanName sql.NullString
	result := middleware.SelectPlan(plan.PlanName, plan.PlanAcronym)
	err := result.Scan(&PlanName)

	if err != sql.ErrNoRows {
		error_message := fmt.Sprintf(`Plan Name "%s" already exists for Application "%s"`, plan.PlanName, plan.PlanAcronym)
		middleware.ErrorHandler(c, 400, error_message)
	} else if err == sql.ErrNoRows {
		_, err := middleware.InsertPlan(plan.PlanName, plan.PlanAcronym, plan.PlanColor, plan.StartDate, plan.EndDate)
		if err != nil {
			log.Fatalln(err)
			return
		}
		c.JSON(200, gin.H{"code": 200, "message": "Plan " + plan.PlanName + " was successfully created!"})
	} else {
		log.Fatalln(err)
		return
	}
	
}