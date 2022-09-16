package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Go struct in the form of JSON
type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	User_group string `json:"user_group"`
	Status     string `json:"status"`
}

var db *sql.DB
var err error

func main() {

	connectionToMySQL()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{code: 200}`)
	})

	http.HandleFunc("/admin-update-user", adminUpdateUserController)

	fmt.Printf("Starting server at port 4000\n")
	err := http.ListenAndServe(":4000", nil)
	checkError(err)

}

func adminUpdateUserController(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	checkError(err)

	// returns a slice of bytes
	body, _ := io.ReadAll(req.Body)

	keyVal := make(map[string]string)
	// func Unmarshal(data []byte, v interface{}) error
	json.Unmarshal(body, &keyVal)

	// Set Headers for response, server informs client that JSON data is being sent.
	w.Header().Set("Content-Type", "application/json")

	username := strings.TrimSpace(keyVal["username"])
	password := keyVal["password"]
	email := strings.TrimSpace(keyVal["email"])
	user_group := keyVal["user_group"]
	status := keyVal["status"]

	adminUpdateUser(username, password, email, user_group, status, w)
}

func adminUpdateUser(username string, password string, email string, user_group string, status string, w http.ResponseWriter) {
	hashedPassword := hashAndSaltPassword([]byte(password))

	if username != "" {
		rows, err := db.Query(`SELECT * FROM accounts WHERE username = ?;`, username)
		checkError(err)
		if rows.Next() {
			adminUpdateUserPassword(username, hashedPassword, email, user_group, status, w)
		} else {
			http.Error(w, "This username does not exist, please try again", 404)
		}
	} else {
		http.Error(w, "Please enter a username.", 500)
	}
}

func adminUpdateUserPassword(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	if hashedPassword != "" {
		adminUpdateUserEmail(username, hashedPassword, email, user_group, status, w)
	} else {
		hashedPassword = getCurrentUserData(username)["password"]
		adminUpdateUserEmail(username, hashedPassword, email, user_group, status, w)
	}
}

func adminUpdateUserEmail(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	if email != "" {
		rows, err := db.Query(`SELECT * FROM accounts WHERE email = ?;`, email)
		checkError(err)
		if rows.Next() {
			http.Error(w, "Email already exists in database. Please try again.", 404)
		} else {
			adminUpdateUserGroup(username, hashedPassword, email, user_group, status, w)
		}
	} else {
		email = getCurrentUserData(username)["email"]
		adminUpdateUserGroup(username, hashedPassword, email, user_group, status, w)
	}
}

func adminUpdateUserGroup(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	if user_group != "" {
		user_group = appendNewUserGroup(username, user_group)
		adminUpdateUserStatus(username, hashedPassword, email, user_group, status, w)
	} else {
		user_group = getCurrentUserData(username)["user_group"]
		adminUpdateUserStatus(username, hashedPassword, email, user_group, status, w)
	}
}

func adminUpdateUserStatus(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	if status != "" {
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, w)
	} else {
		status = getCurrentUserData(username)["status"]
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, w)
	}
}

func adminUpdateAccountsTable(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	_, err := db.Query(`UPDATE accounts SET username = ?, password = ?, email = ?, user_group = ?, status = ? WHERE username = ?`,
		username, hashedPassword, email, user_group, status, username)
	checkError(err)

	jsonStatus := struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}{
		Message: "Successfully updated user!",
		Code:    200,
	}

	json.NewEncoder(w).Encode(jsonStatus)
	// func NewEncoder(w io.Writer) *Encoder
	// func (enc *Encoder) Encode(v any) error
	// Conversion of Go values to JSON
}

func getCurrentUserData(username string) map[string]string {
	var password string
	var email string
	var user_group string
	var status string
	rows, err := db.Query(`SELECT username, password, email, user_group, status FROM accounts WHERE username = ?`,
		username)
	checkError(err)

	currentUserData := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(&username, &password, &email, &user_group, &status)
		checkError(err)
		currentUserData["password"] = password
		currentUserData["email"] = email
		currentUserData["user_group"] = user_group
		currentUserData["status"] = status
	}
	return currentUserData
}

func appendNewUserGroup(username string, user_group string) string {
	currentUserGroup := getCurrentUserData(username)["user_group"]
	currentUserGroupSplit := strings.Split(currentUserGroup, ",")
	newUserGroupSplit := strings.Split(user_group, ",")

	userGroupSlice := []string{}
	for _, i := range newUserGroupSplit {
		if !contains(currentUserGroupSplit, i) {
			updateUserGroupTable(username, i)
			userGroupSlice = append(currentUserGroupSplit, i)
		} else {
			userGroupSlice = currentUserGroupSplit
		}
	}
	user_group = strings.Join(userGroupSlice, ",")
	return user_group
}

func contains(s []string, str string) bool {
	for _, i := range s {
		if i == str {
			return true
		}
	}
	return false
}

func updateUserGroupTable(username string, user_group string) {
	rows, err := db.Query(`SELECT * FROM usergroup WHERE username = ? AND user_group = ?;`,
		username, user_group)
	checkError(err)

	if !rows.Next() {
		fmt.Println("reached here")
		fmt.Println(username, user_group)
		_, err := db.Query(`INSERT INTO usergroup VALUES (?,?)`, username, user_group)
		checkError(err)
	}
}

func hashAndSaltPassword(pwd []byte) string {
	pwdCost := 10
	hash, err := bcrypt.GenerateFromPassword(pwd, pwdCost)
	checkError(err)

	return string(hash)
}

func connectionToMySQL() {
	// db, err := sql.Open(driver, dataSourceName)
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/C3_database")
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Connected to MySQL Database!")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
