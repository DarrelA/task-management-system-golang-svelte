package route

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	// import middleware pkg
	"backend/api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

// Struct tags for key names when serialized into JSON
type User struct {
	LoggedInUser   string   `json:"loggedInUser"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	Email          string   `json:"email"`
	AdminPrivilege int      `json:"admin_privilege"`
	Usergroup      []string `json:"user_group"`
	Status         string   `json:"status"`
	Timestamp      string
}

type ExistingUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	AdminPrivilege string `json:"admin_privilege"`
	Usergroup      string `json:"user_group"`
	Status         string `json:"status"`
	Timestamp      string `json:"created"`
}

func AdminCreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}

	checkGroup := middleware.CheckGroup(newUser.LoggedInUser, "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	whiteSpace := middleware.CheckWhiteSpace(newUser.Username)
	if whiteSpace {
		middleware.ErrorHandler(c, 400, "Username should not contain whitespace")
		return
	}

	minLength := middleware.CheckLength(newUser.Username)
	if minLength {
		middleware.ErrorHandler(c, 400, "Username should not be empty")
		return
	}

	// Validation of password, email
	validPassword := middleware.CheckPassword(newUser.Password)

	// Invalid password format
	if !validPassword {
		middleware.ErrorHandler(c, 400, "Password length should be between length 8 - 10 with numbers and special characters")
		return
	}

	if len(newUser.Email) != 0 {
		emailExist := middleware.CheckEmail(newUser.Email)
		fmt.Print(emailExist)

		// Invalid email format
		if !emailExist {
			middleware.ErrorHandler(c, 400, "Invalid email")
			return
		}
	}

	// Check if username exist before creating
	checkUsername := "SELECT username FROM accounts WHERE username = ?"

	result := db.QueryRow(checkUsername, newUser.Username)

	// Switch between different error case
	switch err := result.Scan(&newUser.Username); err {

	// New user
	case sql.ErrNoRows:

		hash, _ := middleware.GenerateHash(newUser.Password)

		// Convert slice of strings into a single string using strings pkg
		usergroupStr := strings.Join(newUser.Usergroup, ",")

		// Check admin privilege
		if strings.Contains(usergroupStr, "Admin") {
			newUser.AdminPrivilege = 1
		} else {
			newUser.AdminPrivilege = 0
		}

		// INSERT into accounts table
		_, err := db.Exec("INSERT INTO accounts (username, password, email, admin_privilege, user_group, status, timestamp) VALUES (?, ?, ?, ?, ?, ?, NOW())",
			newUser.Username, hash, newUser.Email, newUser.AdminPrivilege, usergroupStr, newUser.Status)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(c, 400, "Invalid field")
			return
		}
		c.JSON(201, gin.H{"code": 201, "message": "New user created"})

		// LOOP through Usergroup slice and validate
		for _, group := range newUser.Usergroup {
			// LOOP to validate group name
			var user_group string

			getGroupname := "SELECT user_group FROM groupnames WHERE user_group = ?"
			result := db.QueryRow(getGroupname, group)

			switch err := result.Scan(&user_group); err {

			// New group name
			case sql.ErrNoRows:
				// INSERT user_group into groupnames table
				_, err := db.Exec("INSERT INTO groupnames (user_group) VALUES (?)", group)
				if err != nil {
					fmt.Println(err)
					middleware.ErrorHandler(c, 400, "Invalid field")
					return
				}
			}

			// Create a composite key for usergroup table
			_, err := db.Exec("INSERT INTO usergroup (username, user_group) VALUES (?, ?)", newUser.Username, group)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Added new composite key %v to %v \n", newUser.Username, group)
		}

	// Username exist
	case nil:
		middleware.ErrorHandler(c, 400, "Existing username")
		return
	}

}

func GetUsers(c *gin.Context) {
	var existingUser ExistingUser
	var data []ExistingUser

	rows, err := db.Query("SELECT username, email, status, admin_privilege, user_group FROM accounts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&existingUser.Username, &existingUser.Email, &existingUser.Status, &existingUser.AdminPrivilege, &existingUser.Usergroup)
		if err != nil {
			panic(err)
		}

		// Using ExistingUser struct as response struct to maintain order of JSON key values
		response := ExistingUser{
			Username:       existingUser.Username,
			Email:          existingUser.Email,
			AdminPrivilege: existingUser.AdminPrivilege,
			Status:         existingUser.Status,
			Usergroup:      existingUser.Usergroup,
			Timestamp:      existingUser.Timestamp,
		}

		// append response into slice
		data = append(data, response)
	}

	// send data as array of JSON obj
	c.JSON(200, data)

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
