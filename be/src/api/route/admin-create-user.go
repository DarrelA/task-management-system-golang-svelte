package route

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// import middleware pkg
	"api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Function to load env file values based on key param
func loadENV(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

// Struct tags for key names when serialized into JSON
type User struct {
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

func adminCreateUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// Start mysql connection with db
	// db, err := sql.Open("mysql", "root:admin123@/c3_database")
	// defer db.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	whiteSpace := middleware.CheckWhiteSpace(newUser.Username)
	if whiteSpace == true {
		middleware.ErrorHandler(c, 400, "Username should not contain whitespace")
		return
	}

	minLength := middleware.CheckLength(newUser.Username)
	if minLength == true {
		middleware.ErrorHandler(c, 400, "Username should not be empty")
		return
	}
	
	// Check if username exist before creating
	checkUsername := "SELECT username FROM accounts WHERE username = ?"

	// Expects a single row result. Returns the first result
	result := db.QueryRow(checkUsername, newUser.Username)

	// Switch between different error case
	switch err := result.Scan(&newUser.Username); err {
	
	// New user
	case sql.ErrNoRows:
		// Validation of password, email
		validPassword := middleware.CheckPassword(newUser.Password)

		// Invalid password format
		if validPassword == false {
			middleware.ErrorHandler(c, 400, "Password length should be between length 8 - 10 with numbers and special characters")
			return
		}

		hash, _ := middleware.GenerateHash(newUser.Password)

		validEmail := middleware.CheckEmail(newUser.Email)

		// Invalid email format
		if validEmail == false {
			middleware.ErrorHandler(c, 400, "Invalid email format")
			return
		}

		// Convert slice of strings into a single string using strings pkg
		usergroupStr := strings.Join(newUser.Usergroup, ", ")

		// INSERT into accounts table
		_, err := db.Exec("INSERT INTO accounts (username, password, email, admin_privilege, user_group, status, timestamp) VALUES (?, ?, ?, ?, ?, ?, NOW())",
			newUser.Username, hash, newUser.Email, newUser.AdminPrivilege, usergroupStr, newUser.Status)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(c, 400, "Invalid field")
			return
		}
		c.JSON(200, gin.H{"code": 200, "message": "New user created"})

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

	default:
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Invalid field")
		return
	}

}

func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT username, email, admin_privilege, user_group, status, timestamp FROM accounts")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var existingUser ExistingUser

		err = rows.Scan(&existingUser.Username, &existingUser.Email, &existingUser.AdminPrivilege, &existingUser.Usergroup, &existingUser.Status, &existingUser.Timestamp)
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

		c.JSON(200, response)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func connectionToDatabase () {
	var err error
	db, err = sql.Open("mysql", "root:admin123@/c3_database")
	if err != nil {
		log.Fatal(err)
		return
	}
	// defer db.Close()
}

// Declare routes
func main() {

	// Start db connection
	connectionToDatabase()
	defer db.Close() 

	router := gin.Default()
	router.POST("/admin-create-user", adminCreateUser)

	// Testing route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "pong",
		})
	})

	router.GET("/get-users", getUsers)

	// using env variables
	port := loadENV("SERVER_PORT")
	server := fmt.Sprintf(":%s", port)

	router.Run(server)
}
