package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

// Function to load env file values based on key param
func loadENV(key string) string {
	err := godotenv.Load("../../../config/jiewei.env")
	checkError(err)
	return os.Getenv(key)
}

func main() {

	connectionToMySQL()
	defer db.Close()

	router := gin.Default()
	router.POST("/admin-update-user", adminUpdateUserController)

	port := loadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}

func adminUpdateUserController(c *gin.Context) {

	var updateUser User

	if err := c.BindJSON(&updateUser); err != nil {
		checkError(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	adminUpdateUser(updateUser.Username, updateUser.Password, updateUser.Email, updateUser.User_group, updateUser.Status, c)
}

func adminUpdateUser(username string, password string, email string, user_group string, status string, c *gin.Context) {

	if username != "" {
		rows, err := db.Query(`SELECT * FROM accounts WHERE username = ?;`, username)
		checkError(err)
		if rows.Next() {
			adminUpdateUserPassword(username, password, email, user_group, status, c)
		} else {
			middleware.ErrorHandler(c, http.StatusNotFound, "Username does not exist. Please try again.")
		}
	} else {
		middleware.ErrorHandler(c, http.StatusNotAcceptable, "Please enter a username")
	}
}

func adminUpdateUserPassword(username string, password string, email string, user_group string, status string, c *gin.Context) {

	if password != "" {
		if middleware.CheckPassword(password) {
			hashedPassword := hashAndSaltPassword([]byte(password))
			adminUpdateUserEmail(username, hashedPassword, email, user_group, status, c)
		} else {
			middleware.ErrorHandler(c, http.StatusBadRequest, "Password length must be between length 8 - 10 with alphabets, numbers and special characters.")
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
				middleware.ErrorHandler(c, http.StatusNotAcceptable, "Email already exists in database. Please try again.")
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
		fmt.Println("1", user_group)
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
	_, err := db.Query(`UPDATE accounts SET username = ?, password = ?, email = ?, user_group = ?, status = ? WHERE username = ?`,
		username, hashedPassword, email, user_group, status, username)
	checkError(err)

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "User was successfully updated!"})
}

func getCurrentUserData(username string) map[string]string {
	var password, email, user_group, status string
	rows, err := db.Query(`SELECT password, email, user_group, status FROM accounts WHERE username = ?`,
		username)
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

func connectionToMySQL() {
	// db, err := sql.Open(driver, dataSourceName)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", loadENV("MYSQL_USERNAME"), loadENV("MYSQL_PASSWORD"), loadENV("MYSQL_PORT"), loadENV("MYSQL_DATABASE"))
	db, err = sql.Open("mysql", dataSourceName)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Connected to MySQL Database!")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}
}
