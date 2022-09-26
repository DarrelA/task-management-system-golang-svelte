package middleware

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Takes in username, user_group to check as params of the function and return a bool
func CheckGroup(username string, usergroup string) bool {

	var (
		checkgroup = false
	)
	queryCheck := "SELECT username, user_group FROM accounts WHERE username = ? AND user_group = ?"

	rows := db.QueryRow(queryCheck, username, usergroup)

	switch err := rows.Scan(&username, &usergroup); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		checkgroup = false

	// Username & Usergroup exists in database
	case nil:
		checkgroup = true
	}

	fmt.Println("iosdahjgiodfg", checkgroup)

	return checkgroup
}

// check if an array contains a given value
// fetch user_group where username = ?
// for loop user_group, append into slice []
// check contains with loop
// func contains(s []string, str string) bool {

// 	var checkGroup = false

// 	for _, i := range s {
// 		if i == str {
// 			checkGroup = true
// 		}
// 	}
// 	return checkGroup
// }
