package route

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserGroup(c *gin.Context) {
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

type usersGroup struct {
	Groupname string `json:"user_group"`
	Usercount int    `json:"user_count"`
}

func GetUsersInGroup(c *gin.Context) {
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
			
		groups +=  "," + usergroup;

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
