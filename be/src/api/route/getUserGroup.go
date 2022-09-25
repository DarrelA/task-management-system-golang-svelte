package route

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserGroup(c *gin.Context) {
	// SELECT * FROM usergroup
	result, err := db.Query("SELECT * FROM groupnames")
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
