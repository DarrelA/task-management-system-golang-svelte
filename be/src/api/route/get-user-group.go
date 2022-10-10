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

	result, err := middleware.SelectUserGroup()
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
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Admin")
	fmt.Println("getstring:", c.GetString("username"))
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	var data []usersGroup
	var groupname string
	var count int
	var usergroup string
	var groups string

	result, err := middleware.SelectUserGroupFromAccountsGroupByUsername()
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	for result.Next() {
		if err := result.Scan(&usergroup); err != nil {
			log.Fatal(err)
		}

		groups += "," + usergroup
	}

	rows, err := middleware.SelectUserGroup()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&groupname); err != nil {
			log.Fatal(err)
		}

		count = strings.Count(groups, groupname)

		response := usersGroup{
			Groupname: groupname,
			Usercount: count,
		}

		data = append(data, response)
	}

	c.JSON(200, data)
}
