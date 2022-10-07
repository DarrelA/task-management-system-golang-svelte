package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.BindJSON(&plan); err != nil {
		return
	}
	// static groupname or use permit open?
	checkPermit := middleware.CheckGroup(c.GetString("username"), "Project manager")
	if checkPermit {
		//proceed with create plan
		if plan.PlanAcronym == "" || plan.PlanName == "" {
			middleware.ErrorHandler(c, 400, "Empty fields!")
			return
		}

		if plan.StartDate == "" || plan.EndDate == "" {
			middleware.ErrorHandler(c, 400, "Start and/or end date is empty!")
			return
		}

		//check plan
		checkWS := middleware.CheckWhiteSpace(plan.PlanName)
		if !checkWS {
			createPlan(plan, c)
		} else {
			middleware.ErrorHandler(c, 400, "Plan name contains whitespace!")
			return
		} 
	} else {
		middleware.ErrorHandler(c, 400, "Unauthorized to create plan")
		return
	}
}

func createPlan(plan models.Plan, c *gin.Context) {
	fmt.Println("plan name ->", plan.PlanName)
	res, err := middleware.SelectPlan(plan.PlanName, plan.PlanAcronym)
	
	if err != nil {
		//middleware.ErrorHandler(c, 400, "Error in select plan")
		log.Fatal(err)
	}
	
	defer res.Close()

	if res.Next() {
		for res.Next() {
			var ep models.Plan
	
			err = res.Scan(&ep.PlanName)
	
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%v\n", ep)
			}
		}
		fmt.Println("plan name exists!")
		middleware.ErrorHandler(c, 400, "Plan name exists!")
	} else {
		//insert plan here
		//fmt.Println("insert plan")
		res, err := middleware.InsertPlan(plan.PlanName, plan.PlanAcronym, plan.PlanColor, plan.StartDate, plan.EndDate)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(res.RowsAffected())
		}

		c.JSON(200, gin.H{"code": 200, "message": "Plan " + plan.PlanName + " created!"})
	}
}