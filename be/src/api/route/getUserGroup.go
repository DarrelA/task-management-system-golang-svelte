package route

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserGroup(c *gin.Context) {
	// SELECT * FROM usergroup
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
	// SELECT * FROM usergroup
	result, err := db.Query("SELECT groupnames.user_group, COUNT(usergroup.username) FROM groupnames LEFT JOIN usergroup ON groupnames.user_group = usergroup.user_group GROUP BY groupnames.user_group")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	var data []usersGroup
	for result.Next() {
		var username int
		var usergroup string

		if err := result.Scan(&usergroup, &username); err != nil {
			log.Fatal(err)
		}

		response := usersGroup{
			Groupname: usergroup,
			Usercount: username,
		}

		data = append(data, response)
	}

	c.JSON(200, data)
}
