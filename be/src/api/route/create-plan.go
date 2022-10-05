package route

import (
	"backend/api/middleware"
	"backend/api/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreatePlan(c *gin.Context) {
	var np models.Plan
	if err := c.BindJSON(&np); err != nil {
		return
	}
	// static groupname, to change to dynamic - permit open
	//checkPermit := middleware.CheckGroup(c.GetString("username"), "Project manager")
	//if checkPermit {
		//proceed with create plan
		if np.PlanAcronym != "" {
			if np.PlanName != "" {
				//checkPlan
				checkWS := middleware.CheckWhiteSpace(np.PlanName)
				if !checkWS {
					createPlan(np, c)
				} else {
					middleware.ErrorHandler(c, 400, "Plan name contains whitespace!")
				} 
			} else {
				middleware.ErrorHandler(c, 400, "Empty plan name!")	
			}
		} else {
			middleware.ErrorHandler(c, 400, "Empty app name!")	
		}
	// } else {
	// 	middleware.ErrorHandler(c, 400, "Unauthorized access")
	// }
}

func createPlan(np models.Plan, c *gin.Context) {
	fmt.Println("plan name ->", np.PlanName)
	res, err := middleware.SelectPlan(np.PlanName, np.PlanAcronym)
	
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
		middleware.ErrorHandler(c, 400, "plan name exists!")
	} else {
		//insert plan here
		fmt.Println("insert plan")
		res, err := middleware.InsertPlan(np.PlanName, np.PlanAcronym, np.PlanColor, np.StartDate, np.EndDate)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(res.RowsAffected())
		}

		c.JSON(200, gin.H{"code": 200, "message": "test create plan"})
	}
}