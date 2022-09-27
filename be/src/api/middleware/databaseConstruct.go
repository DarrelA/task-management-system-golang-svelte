package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var (
	queryInsertAccounts = "INSERT INTO accounts (username, password, email, admin_privilege, user_group, status, timestamp) VALUES (?,?,?,?,?,?,now());"

	queryInsertUserGroup  = "INSERT INTO usergroup (username, user_group) VALUES (?,?);"

	queryInsertGroupnames = "INSERT INTO groupnames (user_group) VALUES (?);"
)

var (
	querySelectAccounts           = "SELECT username, email, user_group, status FROM accounts;"

	querySelectAccountByLogin     = "SELECT username, password, status FROM accounts WHERE username = ?;"

	querySelectAccountsByUsername = "SELECT username, password, email, admin_privilege, user_group, status, timestamp FROM accounts WHERE username = ?;"

	querySelectUserGroupByUsernameUserGroup = "SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?;"

	querySelectGroupnamesByUserGroup = "SELECT user_group FROM groupnames WHERE user_group = ?;"

	querySelectCompositeKey = "SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?"

	querySelectCheckGroupFromAccounts = "SELECT username, user_group FROM accounts WHERE username = ?"

	querySelectUserFromUserGroupByUsername = "SELECT user_group FROM accounts WHERE username = ?"

	querySelectUserGroupFromGroupnamesByUserGroup = "SELECT user_group FROM groupnames WHERE user_group = ?"

	querySelectUsernameFromAccountsByUsername = "SELECT username FROM accounts WHERE username = ?"

	querySelectUserGroup = "SELECT user_group FROM groupnames"

	querySelectUserGroupFromAccountsGroupByUsername = "SELECT user_group FROM accounts GROUP BY username"

	querySelectPasswordEmailFromAccountsByUsername = "SELECT password, email FROM accounts WHERE username = ?"
)

var (
	queryUpdateAccountsAdmin = "UPDATE accounts SET password = ?, email = ?, admin_privilege = ?, user_group = ?, status = ? WHERE username = ?;"

	queryUpdateAccountsSetUsernameByUserGroup = "UPDATE accounts SET user_group = ? WHERE username = ?"

	queryUpdateUserToDb = "UPDATE accounts SET password = ?, email = ? WHERE username = ?"
)

// INSERT
func InsertNewAccount(username string, password string, email string, admin_privilege int, user_group string, status string) (sql.Result, error) {
	result, err := db.Exec(queryInsertAccounts, username, password, email, admin_privilege, user_group, status)
	return result, err
}

func InsertUserGroup(username string, user_group string) (sql.Result, error) {
	result, err := db.Exec(queryInsertUserGroup, username, user_group)
	return result, err
}

func InsertGroupnames(user_group string) (sql.Result, error) {
	result, err := db.Exec(queryInsertGroupnames, user_group)
	return result, err
}

// SELECT
func SelectUserGroup() (*sql.Rows, error) {
	result, err := db.Query(querySelectUserGroup)
	return result, err
}

// For CheckGroup //
func SelectCheckGroupFromAccounts(username string) *sql.Row {
	result := db.QueryRow(querySelectCheckGroupFromAccounts, username)
	return result
}


func SelectPasswordEmailFromAccountsByUsername(username string) (*sql.Rows, error) {
	result, err := db.Query(querySelectPasswordEmailFromAccountsByUsername, username)
	return result, err
}

func SelectUserGroupFromAccountsGroupByUsername() (*sql.Rows, error) {
	result, err := db.Query(querySelectUserGroupFromAccountsGroupByUsername)
	return result, err
}

func SelectUserFromUserGroupByUsername(username string) *sql.Row {
	result := db.QueryRow(querySelectUserFromUserGroupByUsername, username)
	return result
}

func SelectUserGroupFromGroupnamesByUserGroup(user_group string) *sql.Row {
	result := db.QueryRow(querySelectUserGroupFromGroupnamesByUserGroup, user_group)
	return result
}

func SelectUsernameFromAccountsByUsername(username string) *sql.Row {
	result := db.QueryRow(querySelectUsernameFromAccountsByUsername, username)
	return result
}

func SelectAccounts() (*sql.Rows, error) {
	result, err := db.Query(querySelectAccounts)
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

func SelectGroupnamesbyUserGroup(user_group string) *sql.Row {
	result := db.QueryRow(querySelectGroupnamesByUserGroup, user_group)
	return result
}

func SelectCompositeKey(username string, user_group string) *sql.Row {
	result := db.QueryRow(querySelectCompositeKey, username, user_group)
	return result
}

// UPDATE
func UpdateAccountsAdmin(password string, email string, admin_privilege int, user_group string, status string, username string, c *gin.Context) (*sql.Rows, error) {
	result, err := db.Query(queryUpdateAccountsAdmin, password, email, admin_privilege, user_group, status, username)
	return result, err
}

func UpdateAccountsSetUsernameByUsergroup(user_group string, username string) (*sql.Rows, error) {
	result, err := db.Query(queryUpdateAccountsSetUsernameByUserGroup, user_group, username)
	return result, err
}

func UpdateUserToDb(password string, email string, username string) (sql.Result, error) {
	 result, err := db.Exec(queryUpdateUserToDb, password, email, username)
	return result, err
}
