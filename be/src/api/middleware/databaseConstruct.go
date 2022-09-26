package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var (
	queryInsertAccounts = "INSERT INTO accounts (username, password, email, admin_privilege, user_group, status, timestamp) VALUES (?,?,?,?,?,?,now());"

	queryInsertUserGroup  = "INSERT INTO usergroup (username, user_group) VALUES (?,?);"
	queryInsertGroupnames = "INSERT INTO groupnames (user_group) VALUES (?,?);"
)

var (
	querySelectAccounts           = "SELECT username, password, email, admin_privilege, user_group, status, timestamp FROM accounts;"
	querySelectAccountByLogin     = "SELECT username, password, status FROM accounts WHERE username = ?;"
	querySelectAccountsByUsername = "SELECT username, password, email, admin_privilege, user_group, status, timestamp FROM accounts WHERE username = ?;"

	querySelectUserGroup                    = "SELECT user_group FROM usergroup;"
	querySelectUserGroupByUsernameUserGroup = "SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?;"

	querySelectGroupnamesByUserGroup = "SELECT user_group FROM groupnames WHERE user_group = ?;"
)

var (
	queryUpdateAccountsAdmin = "UPDATE accounts SET password = ?, email = ?, admin_privilege = ?, user_group = ?, status = ? WHERE username = ?;"
)

func InsertUserGroup(username string, user_group string) (sql.Result, error) {
	result, err := db.Exec(queryInsertUserGroup, username, user_group)
	return result, err
}

func SelectAccountByLogin(username string, c *gin.Context) *sql.Row {
	result := db.QueryRow(querySelectAccountByLogin, username)
	return result
}

func SelectAccountsByUsername(username string, c *gin.Context) *sql.Row {
	result := db.QueryRow(querySelectAccountsByUsername, username)
	return result
}

func SelectUserGroupByUsernameUserGroup(username string, user_group string) *sql.Row {
	result := db.QueryRow(querySelectUserGroupByUsernameUserGroup, username, user_group)
	return result
}

func UpdateAccountsAdmin(password string, email string, admin_privilege int, user_group string, status string, username string, c *gin.Context) (*sql.Rows, error) {
	result, err := db.Query(queryUpdateAccountsAdmin, password, email, admin_privilege, user_group, status, username)
	return result, err
}
