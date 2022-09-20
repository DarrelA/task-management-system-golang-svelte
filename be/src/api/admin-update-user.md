# Admin Update User

### Accounts Table

1. check if `username` is provided
1. retrieve username, password, email, usergroup, status (isactive) from frontend.
1. trim username and email
1. check if username exist (otherwise return error)
```
SELECT * FROM accounts WHERE username = ?
```
1. 


```go
if (username != "") {
    // Check if username exists in database
    // Error: Username does not exists in database. Please try again.
    if (password != "") {
        // update accounts with NEW password provided
        if (email != "") {
            // Check if email exists in database
            // Error: Email already exists in database. Please try again.
            // Update Accounts with NEW email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        } else {
            // update accounts with OLD email, SELECT statement to get old email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        }
    } else {
        // update accounts with OLD password, SELECT statement to get old password
        if (email != "") {
            // Check if email exists in database
            // Error: Email already exists in database. Please try again.
            // Update Accounts with NEW email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        } else {
            // update accounts with OLD email, SELECT statement to get old email
            if (user_group) {
                // Append new user_group to accounts table (if any)
                // Add new usergroup and username to the `user_group` table (if any)
                // Update `accounts` table with NEW user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            } else {
                // update accounts with OLD user_group, SELECT statement to get old user_group
                if (status) {
                    // Update Accounts with NEW status
                } else {
                    // update accounts with OLD status, SELECT statement to get old status
                }
            }
        }
    }
} else {
    // Error: Please provide a username.
}

```

- Calling Backend API using `http.HandleFunc`

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"

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

	fmt.Println(keyVal)

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

	if username != "" {
		rows, err := db.Query(`SELECT * FROM accounts WHERE username = ?;`, username)
		checkError(err)
		if rows.Next() {
			adminUpdateUserPassword(username, password, email, user_group, status, w)
		} else {
			responseMessage("Username does not exist. Please try again.", 404, w)
		}
	} else {
		responseMessage("Please enter a username.", 500, w)
	}
}

func adminUpdateUserPassword(username string, password string, email string, user_group string, status string, w http.ResponseWriter) {

	if password != "" {
		if validatePassword(password, w) {
			hashedPassword := hashAndSaltPassword([]byte(password))
			adminUpdateUserEmail(username, hashedPassword, email, user_group, status, w)
		} else {
			responseMessage("Password length must be between length 8 - 10 with alphabets, numbers and special characters.", 400, w)
		}
	} else {
		password = getCurrentUserData(username)["password"]
		adminUpdateUserEmail(username, password, email, user_group, status, w)
	}

}

func adminUpdateUserEmail(username string, hashedPassword string, email string, user_group string, status string, w http.ResponseWriter) {
	if email != "" {
		rows, err := db.Query(`SELECT * FROM accounts WHERE email = ?;`, email)
		checkError(err)
		if rows.Next() {
			responseMessage("Email already exists in database. Please try again.", 500, w)
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
		fmt.Println("1", user_group)
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

	responseMessage("User successfully updated!", 200, w)
}

func getCurrentUserData(username string) map[string]string {
	var password, email, user_group, status string
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

	userGroupSlice := []string{currentUserGroup}
	for _, i := range newUserGroupSplit {
		if !contains(currentUserGroupSplit, i) {
			updateUserGroupTable(username, i)
			userGroupSlice = append(userGroupSlice, i)
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
		_, err := db.Query(`INSERT INTO usergroup VALUES (?,?)`, username, user_group)
		checkError(err)
	}
}

func responseMessage(Message string, Code int, w http.ResponseWriter) bool {
	// func NewEncoder(w io.Writer) *Encoder
	// func (enc *Encoder) Encode(v any) error
	// Conversion of Go values to JSON

	jsonStatus := struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}{
		Message: Message,
		Code:    Code,
	}

	bs, err := json.Marshal(jsonStatus)
	checkError(err)
	io.WriteString(w, string(bs))
	// json.NewEncoder(w).Encode(jsonStatus)
	return false
}

func validatePassword(password string, w http.ResponseWriter) bool {
	var (
		hasMinLength = false
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		hasSpecial   = false
	)

	if len(password) >= 8 && len(password) <= 10 {
		hasMinLength = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSpecial = true
		}
	}
	return hasMinLength && hasUpper && hasLower && hasNumber && hasSpecial
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

```

