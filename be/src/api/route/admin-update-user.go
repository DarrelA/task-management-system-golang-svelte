package route

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"backend/api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Go struct in the form of JSON
type UserData struct {
	LoggedInUser string `json:"loggedInUser"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	User_group   string `json:"user_group"`
	Status       string `json:"status"`
}

type SpecificUser struct {
	Username string `json:"username"`
}

func AdminUpdateUser(c *gin.Context) {

	var updateUser UserData

	if err := c.BindJSON(&updateUser); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// Check user group
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	adminUpdateUser(updateUser.Username, updateUser.Password, updateUser.Email, updateUser.User_group, updateUser.Status, c)
}

func adminUpdateUser(username string, password string, email string, user_group string, status string, c *gin.Context) {
	if username != "" {
		adminUpdateUserPassword(username, password, email, user_group, status, c)
	} else {
		middleware.ErrorHandler(c, 400, "Please enter a username")
	}
}

func adminUpdateUserPassword(username string, password string, email string, user_group string, status string, c *gin.Context) {

	if password != "" {
		if middleware.CheckPassword(password) {
			hashedPassword, _ := middleware.GenerateHash(password)
			adminUpdateUserEmail(username, hashedPassword, email, user_group, status, c)
		} else {
			middleware.ErrorHandler(c, 400, "Password length must be between length 8 - 10 with alphabets, numbers and special characters.")
		}
	} else {
		password = getCurrentUserData(username, c)["password"]
		adminUpdateUserEmail(username, password, email, user_group, status, c)
	}

}

func adminUpdateUserEmail(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	if email != "" {
		validEmail := middleware.CheckEmail(email)
		if !validEmail {
			middleware.ErrorHandler(c, 400, "Invalid Email Format")
			return
		}
		whiteSpace := middleware.CheckWhiteSpace(email)
		if whiteSpace {
			fmt.Println("Reached here!")
			middleware.ErrorHandler(c, 400, "Email should not contain whitespace")
			return
		}
		fmt.Println("Here")
		adminUpdateUserGroup(username, hashedPassword, email, user_group, status, c)
	} else {
		email = getCurrentUserData(username, c)["email"]
		adminUpdateUserGroup(username, hashedPassword, email, user_group, status, c)
	}
}

func adminUpdateUserGroup(username string, hashedPassword string, email string, user_group string, status string, c *gin.Context) {
	admin_privilege := getAdminPrivilege(user_group)
	updateUserGroupTable(username, user_group)
	adminUpdateUserStatus(username, hashedPassword, email, user_group, status, admin_privilege, c)
}

func getAdminPrivilege(user_group string) int {
	if strings.Contains(user_group, "Admin") {
		return 1
	} else {
		return 0
	}
}

func adminUpdateUserStatus(username string, hashedPassword string, email string, user_group string, status string, admin_privilege int, c *gin.Context) {
	if status != "" {
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, admin_privilege, c)
	} else {
		status = getCurrentUserData(username, c)["status"]
		adminUpdateAccountsTable(username, hashedPassword, email, user_group, status, admin_privilege, c)
	}
}

func adminUpdateAccountsTable(username string, hashedPassword string, email string, user_group string, status string, admin_privilege int, c *gin.Context) {
	_, err := middleware.UpdateAccountsAdmin(hashedPassword, email, admin_privilege, user_group, status, username, c)
	checkError(err)
	successMessage := fmt.Sprintf("User %s was successfully updated!", username)
	c.JSON(http.StatusCreated, gin.H{"code": 200, "message": successMessage})
}

func getCurrentUserData(username string, c *gin.Context) map[string]string {
	var password, email, user_group, status, timestamp string
	var admin_privilege int
	result := middleware.SelectAccountsByUsername(username, c)

	currentUserData := make(map[string]string)
	err := result.Scan(&username, &password, &email, &user_group, &admin_privilege, &status, &timestamp)
	if err != sql.ErrNoRows {
		currentUserData["password"] = password
		currentUserData["email"] = email
		currentUserData["user_group"] = user_group
		currentUserData["status"] = status
	} else if err != nil {
		checkError(err)
	}
	return currentUserData
}

func updateUserGroupTable(username string, user_group string) {
	if user_group != "" {
		userGroupSlice := strings.Split(user_group, ",")
		for _, value := range userGroupSlice {
			result := middleware.SelectUserGroupByUsernameUserGroup(username, value)
			err := result.Scan(&username, &value)

			if err == sql.ErrNoRows {
				middleware.InsertUserGroup(username, value)
			} else if err != nil {
				checkError(err)
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}
}
