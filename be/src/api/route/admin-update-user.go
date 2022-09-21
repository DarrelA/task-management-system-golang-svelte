package route

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"backend/api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Go struct in the form of JSON
type UpdateUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	User_group string `json:"user_group"`
	Status     string `json:"status"`
}

type SpecificUser struct {
	Username string `json:"username"`
}

// var err error

func AdminUpdateUser(c *gin.Context) {

	var updateUser UpdateUser

	if err := c.BindJSON(&updateUser); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	adminUpdateUser(updateUser.Username, updateUser.Password, updateUser.Email, updateUser.User_group, updateUser.Status, c)
}

func adminUpdateUser(username string, password string, email string, user_group string, status string, c *gin.Context) {
	if username != "" {
		fmt.Println("0")
		rows, err := db.Query(`SELECT * FROM accounts WHERE username = ?;`, username)
		checkError(err)
		if rows.Next() {
			adminUpdateUserPassword(username, password, email, user_group, status, c)
		} else {
			middleware.ErrorHandler(c, 200, "Username does not exist. Please try again.")
		}
	} else {
		middleware.ErrorHandler(c, 200, "Please enter a username")
	}
}

func adminUpdateUserPassword(username string, password string, email string, user_group string, status string, c *gin.Context) {

	if password != "" {
		if middleware.CheckPassword(password) {
			hashedPassword := hashAndSaltPassword([]byte(password))
			adminUpdateUserEmail(username, hashedPassword, email, user_group, status, c)
		} else {
			middleware.ErrorHandler(c, 200, "Password length must be between length 8 - 10 with alphabets, numbers and special characters.")
		}
	} else {
		password = getCurrentUserData(username)["password"]
		adminUpdateUserEmail(username, password, email, user_group, status, c)
	}

}

func adminUpdateUserEmail(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	if email != "" {
		currentEmail := getCurrentUserData(username)["email"]
		if email == currentEmail {
			adminUpdateUserGroup(username, hashedPassword, currentEmail, user_group, status, c)
		} else {
			rows, err := db.Query(`SELECT * FROM accounts WHERE email = ?;`, email)
			checkError(err)
			if rows.Next() {
				middleware.ErrorHandler(c, 200, "Email already exists in database. Please try again.")
			} else {
				adminUpdateUserGroup(username, hashedPassword, email, user_group, status, c)
			}
		}
	} else {
		email = getCurrentUserData(username)["email"]
		adminUpdateUserGroup(username, hashedPassword, email, user_group, status, c)
	}
}

func adminUpdateUserGroup(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	if user_group != "" {
		user_group = appendNewUserGroup(username, user_group)
		adminUpdateUserStatus(username, hashedPassword, email, user_group, status, c)
	} else {
		user_group = getCurrentUserData(username)["user_group"]
		adminUpdateUserStatus(username, hashedPassword, email, user_group, status, c)
	}
}

func adminUpdateUserStatus(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	if status != "" {
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, c)
	} else {
		status = getCurrentUserData(username)["status"]
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, c)
	}
}

func adminUpdateAccountsTable(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	_, err := db.Query(`UPDATE accounts SET password = ?, email = ?, user_group = ?, status = ? WHERE username = ?`,
		hashedPassword, email, user_group, status, username)
	checkError(err)
	successMessage := fmt.Sprintf("User %s was successfully updated!", username)
	c.JSON(http.StatusCreated, gin.H{"code": 201, "message": successMessage})
}

func getCurrentUserData(username string) map[string]string {
	var password, email, user_group, status string
	rows, err := db.Query(`SELECT password, email, user_group, status FROM accounts WHERE username = ?`,
		&username)
	checkError(err)

	currentUserData := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(&password, &email, &user_group, &status)
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

func hashAndSaltPassword(pwd []byte) string {
	pwdCost := 10
	hash, err := bcrypt.GenerateFromPassword(pwd, pwdCost)
	checkError(err)

	return string(hash)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}
}

func GetSelectedUser(c *gin.Context) {
	var specificUser SpecificUser

	if err := c.BindJSON(&specificUser); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	fmt.Println(specificUser.Username)

	var existingUser ExistingUser
	var data []ExistingUser

	rows, err := db.Query("SELECT email, status, user_group FROM accounts WHERE username = ?")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&existingUser.Email, &existingUser.Status, &existingUser.Usergroup, &specificUser.Username,)
		if err != nil {
			panic(err)
		}

		// Using ExistingUser struct as response struct to maintain order of JSON key values
		response := ExistingUser{
			Username:       existingUser.Username,
			Email:          existingUser.Email,
			Status:         existingUser.Status,
			Usergroup:      existingUser.Usergroup,
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
