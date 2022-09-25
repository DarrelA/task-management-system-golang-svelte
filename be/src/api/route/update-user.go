package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "user"
	password = "test1234"
	hostname = "127.0.0.1:3306"
	database = "c3_database"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UpdateUser(c *gin.Context) {
	var u user
	if err := c.BindJSON(&u); err != nil {
		return
	}

	currentData := getSelect(u.Username)

	if u.Password != "" && u.Email == "" {
		newPassword := prePassword(u.Username, u.Password, c)
		updateToDB(u.Username, newPassword, currentData["email"], c)
	} else if u.Password == "" && u.Email != "" {
		check := preEmail(u.Username, u.Email, c)
		if check {
			//fmt.Println("email check ok!")
			updateToDB(u.Username, currentData["password"], u.Email, c)
		} else {
			middleware.ErrorHandler(c, 400, "Invalid email format")
		}
	} else if u.Password != "" && u.Email != "" {
		newPassword := prePassword(u.Username, u.Password, c)
		check := preEmail(u.Username, u.Email, c)
		if check {
			updateToDB(u.Username, newPassword, u.Email, c)
		} else {
			middleware.ErrorHandler(c, 400, "Invalid email format")
		}
	} else { // extra handler
		middleware.ErrorHandler(c, 400, "Empty fields!")
	}
}

func dsn() string {
	// username:password@tcp(127.0.0.1:3306)/database-name")
	//return fmt.Sprintf("%s:%s@/%s", LoadENV("SERVER_USER"), LoadENV("SERVER_PASSWORD"), LoadENV("SERVER_DB"))
	return fmt.Sprintf("%s:%s@/%s", username, password, database)
}

func getSelect(username string) map[string]string {
	db, err := sql.Open("mysql", dsn())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT password, email FROM accounts WHERE username = ?", username)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	currentUserData := make(map[string]string)
	for res.Next() {
		var u user

		err := res.Scan(&u.Password, &u.Email)

		if err != nil {
			log.Fatal(err)
		} else {
			currentUserData["password"] = u.Password
			currentUserData["email"] = u.Email
		}
		fmt.Printf("%v\n", u)
	}
	return currentUserData
}

func hashPassword(password string) string {
	hashedPassword, err := middleware.GenerateHash(password)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func prePassword(username string, password string, c *gin.Context) string {
	var hPassword string
	if username != "" && password != "" {
		check := middleware.CheckPassword(password)
		if check {
			hPassword = hashPassword(password)
		} else {
			middleware.ErrorHandler(c, 400, "Password requirement not met!")
		}
	} else {
		middleware.ErrorHandler(c, 400, "Field(s) is empty")
	}
	return hPassword
}

func preEmail(username string, email string, c *gin.Context) bool {
	var check bool
	if username != "" && email != "" {
		check = middleware.CheckEmail(email)
	} else {
		middleware.ErrorHandler(c, 400, "Field(s) is empty")
	}
	return check
}

func updateToDB(username string, newPassword string, newEmail string, c *gin.Context) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}

	//update statement
	stmt, err := db.Prepare("UPDATE accounts SET password = ?, email = ? WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}

	//execute
	res, err := stmt.Exec(newPassword, newEmail, username)
	if err != nil {
		log.Fatal(err)
	}

	a, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows affected: ", a)
	middleware.ErrorHandler(c, 200, "Update success!")
}
