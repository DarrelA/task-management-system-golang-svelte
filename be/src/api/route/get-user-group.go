package route

import (
	"backend/api/middleware"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type usersGroup struct {
	LoggedInUser string `json:"loggedInUser"`
	Groupname    string `json:"user_group"`
	Usercount    int    `json:"user_count"`
}

func GetUserGroup(c *gin.Context) {
	var usersGroups usersGroup
	checkGroup := middleware.CheckGroup(usersGroups.LoggedInUser, "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	result, err := db.Query("SELECT user_group FROM groupnames")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	var data []string
	for result.Next() {
		var usergroup string

		if err := result.Scan(&usergroup); err != nil {
			log.Fatal(err)
		}
		data = append(data, usergroup)
	}

	c.JSON(200, data)
}

func GetUsersInGroup(c *gin.Context) {
	var usersGroups usersGroup
	// Check user group
	checkGroup := middleware.CheckGroup(usersGroups.LoggedInUser, "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	var data []usersGroup
	var groupname string
	var count int
	var usergroup string
	var groups string

	result, err := db.Query("SELECT user_group FROM accounts GROUP BY username")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	for result.Next() {
		if err := result.Scan(&usergroup); err != nil {
			log.Fatal(err)
		}

		groups += "," + usergroup

		fmt.Println(groups)
	}

	rows, err := db.Query("SELECT user_group FROM groupnames")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&groupname); err != nil {
			log.Fatal(err)
		}

		count = strings.Count(groups, groupname)
		fmt.Println(count)

		response := usersGroup{
			Groupname: groupname,
			Usercount: count,
		}

		data = append(data, response)
	}

	c.JSON(200, data)
}
