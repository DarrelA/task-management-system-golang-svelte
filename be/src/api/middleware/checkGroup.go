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