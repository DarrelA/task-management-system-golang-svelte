package middleware

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Takes in username, user_group to check as params of the function and return a bool
func CheckGroup(username string, usergroup string) bool {

	var (
		checkgroup = false
	)
	queryCheck := "SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?"

	rows := db.QueryRow(queryCheck, username, usergroup)

	switch err := rows.Scan(&username, &usergroup); err {

	// Username/Usergroup does not exist in database
	case sql.ErrNoRows:
		checkgroup = false

	// Username & Usergroup exists in database
	case nil:
		checkgroup = true
	}

	return checkgroup
}

// func main() {
// 	c1 := CheckGroup("admin", "Admin")
// 	c2 := CheckGroup("admin", "Team Member")
// 	fmt.Println(c1)
// 	fmt.Println(c2)
// }

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
