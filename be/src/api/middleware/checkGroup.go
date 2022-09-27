package middleware

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Takes in username, user_group to check as params of the function and return a bool
func CheckGroup(username string, groupname string) bool {
	var user_group string

	var (
		checkgroup = false
	)

	result := SelectCheckGroupFromAccounts(username)

	switch err := result.Scan(&username, &user_group); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		checkgroup = false

	// Username & Usergroup exists in database
	case nil:
		if strings.Contains(user_group, groupname) {
			checkgroup = true
		} else {
			checkgroup = false
		}
	}

	return checkgroup
}
