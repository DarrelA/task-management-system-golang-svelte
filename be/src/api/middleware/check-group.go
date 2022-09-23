package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


// Struct tags for key names when serialized into JSON
type Group struct {
	Username	string	`json:"username"`
	Usergroup	string	`json:"user_group"`
}

func CheckUserGroupController(c *gin.Context) {
	var newGroup Group

	if err := c.BindJSON(&newGroup); err != nil {
		fmt.Println(err)
		ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// check whitespace in username
	whiteSpace := CheckWhiteSpace(newGroup.Username)
	if whiteSpace{
		ErrorHandler(c, 400, "Username should not contain whitespace")
		return
	}

	// check if username empty
	minLength := CheckLength(newGroup.Username)
	if minLength{
		ErrorHandler(c, 400, "Username should not be empty")
		return
	}

	minLengths := CheckLength(newGroup.Usergroup)
	if minLengths{
		ErrorHandler(c, 400, "Usergroup should not be empty")
		return
	}

	getUserGroupData(newGroup.Username, newGroup.Usergroup, c)

}

func getUserGroupData(username string, user_group string, c *gin.Context) {

	var newGroup Group

	db, err := sql.Open("mysql", "root:password@/c3_database")
	//if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	rows := db.QueryRow(`SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?;`, username, user_group)
	if err != nil {
		log.Fatal(err)
	}

	switch err := rows.Scan(&newGroup.Username, &newGroup.Usergroup); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		// error handling
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(contains([]string{user_group}, newGroup.Usergroup ))
		ErrorHandler(c, 400, "Username, Usergroup does not exist")


	// Username & Usergroup exists in database
	case nil: 
		if newGroup.Usergroup == user_group {
			fmt.Println(contains([]string{user_group}, newGroup.Usergroup ))
			User := fmt.Sprintf("User is %s", user_group)
			c.JSON(200, gin.H{"code": 200, "message": User})

			log.Printf(newGroup.Username)
			return
		} 
	}
}

// check if an array contains a given value
// fetch user_group where username = ?
// for loop user_group, append into slice []
// check contains with loop
func contains(s []string, str string) bool {

	var checkGroup = false

	for _, i := range s {
			if i == str {
				checkGroup = true
			}	
	}
	return checkGroup
}