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
	Email          string   `json:"email" validate:"omitempty,email"`
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

	// Check user group
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	// Username Validations
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

	// Password Validations
	validPassword := middleware.CheckPassword(newUser.Password)

	if !validPassword {
		middleware.ErrorHandler(c, 400, "Password length should be between length 8 - 10 with numbers and special characters")
		return
	}

	// Email Validations
	if len(newUser.Email) != 0 {

		// email entered
		checkEmail := middleware.CheckEmail(newUser.Email)

		// Invalid email format
		if !checkEmail {
			fmt.Println("Invalid email")
			middleware.ErrorHandler(c, 400, "Invalid email")
			return
		}
	}

	// Check if username exist before creating
	result := middleware.SelectUsernameFromAccountsByUsername(newUser.Username)

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
		_, err := middleware.InsertNewAccount(newUser.Username, hash, newUser.Email, newUser.AdminPrivilege, usergroupStr, newUser.Status)

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

			result := middleware.SelectGroupnamesbyUserGroup(user_group)

			switch err := result.Scan(&user_group); err {

			// New group name
			case sql.ErrNoRows:
				// INSERT user_group into groupnames table
				_, err := middleware.InsertGroupnames(group)
				if err != nil {
					middleware.ErrorHandler(c, 400, "Invalid field")
					return
				}
			}

			// Create a composite key for usergroup table
			_, err := middleware.InsertUserGroup(newUser.Username, group)
			if err != nil {
				fmt.Println(err)
				return
			}
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

	checkGroup := middleware.CheckGroup(c.GetString("username"), "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	rows, err := middleware.SelectAccounts()
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&existingUser.Username, &existingUser.Email, &existingUser.Usergroup, &existingUser.Status)
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
