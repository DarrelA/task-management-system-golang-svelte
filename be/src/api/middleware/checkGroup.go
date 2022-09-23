package middleware

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Struct tags for key names when serialized into JSON
type Group struct {
	Username  string `json:"username"`
	Usergroup string `json:"user_group"`
}

// Takes in username, user_group to check as params of the function and return a bool
func CheckGroup(c *gin.Context, username string, user_group string) bool {

	var newGroup Group
	var checkgroup = false

	rows := db.QueryRow(`SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?;`, username, user_group)

	switch err := rows.Scan(username, user_group); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		fmt.Println(contains([]string{user_group}, newGroup.Usergroup))
		checkgroup = false

	// Username & Usergroup exists in database
	case nil:
		if newGroup.Usergroup == user_group {
			fmt.Println(contains([]string{user_group}, newGroup.Usergroup))
			checkgroup = true
		}
	}
	return checkgroup
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
