package middleware

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Takes in username, user_group to check as params of the function and return a bool
func CheckGroup(username string, user_group string) bool {

	var (
		checkgroup = false
	)
	
	result := SelectCheckGroupFromAccounts(username, user_group)

	switch err := result.Scan(&username, &user_group); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		checkgroup = false

	// Username & Usergroup exists in database
	case nil:
		checkgroup = true
	}

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
