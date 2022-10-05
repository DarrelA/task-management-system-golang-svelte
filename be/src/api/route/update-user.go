package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// const (
//  username = "user"
//  password = "test1234"
//  hostname = "127.0.0.1:3306"
//  database = "c3_database"
// )

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

	u.Username = c.GetString("username")
	fmt.Println("Username", u.Username)

	currentData := getSelect(u.Username)

	if u.Password != "" && u.Email == "" {
		// newPassword := prePassword(u.Username, u.Password, c)
		check := middleware.CheckPassword(u.Password)
		if !check {
			middleware.ErrorHandler(c, 400, "Password requirement not met!")
			return
		}
		newPassword := hashPassword(u.Password)
		updateToDB(u.Username, newPassword, currentData["email"], c)
		return
	} else if u.Password == "" && u.Email != "" {
		check := preEmail(u.Username, u.Email, c)
		if check {
			updateToDB(u.Username, currentData["password"], u.Email, c)
		} else {
			middleware.ErrorHandler(c, 400, "Invalid email format")
			return
		}
	} else if u.Password != "" && u.Email != "" {
		checkPassword := middleware.CheckPassword(u.Password)
		if !checkPassword {
			middleware.ErrorHandler(c, 400, "Password requirement not met!")
			return
		}
		newPassword := hashPassword(u.Password)
		// newPassword := prePassword(u.Username, u.Password, c)
		checkEmail := preEmail(u.Username, u.Email, c)
		if checkEmail {
			updateToDB(u.Username, newPassword, u.Email, c)
		} else {
			middleware.ErrorHandler(c, 400, "Invalid email format")
			return
		}
	} else { // extra handler
		middleware.ErrorHandler(c, 400, "Empty fields!")
	}
}

// func prePassword(username string, password string, c *gin.Context) string {
//  var hPassword string
//  if username != "" && password != "" {
//    check := middleware.CheckPassword(password)
//    if check {
//      hPassword = hashPassword(password)
//    } else {
//      middleware.ErrorHandler(c, 400, "Password requirement not met!")
//    }
//  }
//  return hPassword
// }

func dsn() string {
	// username:password@tcp(127.0.0.1:3306)/database-name")
	return fmt.Sprintf("%s:%s@/%s", middleware.LoadENV("SERVER_USER"), middleware.LoadENV("SERVER_PASSWORD"), middleware.LoadENV("SERVER_DB"))
	//return fmt.Sprintf("%s:%s@/%s", username, password, database)
}

func getSelect(username string) map[string]string {
	db, err := sql.Open("mysql", dsn())

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	res, err := middleware.SelectPasswordEmailFromAccountsByUsername(username)

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
		//fmt.Printf("%v\n", u)
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

func preEmail(username string, email string, c *gin.Context) bool {
	var check bool
	if username != "" && email != "" {
		whiteSpace := middleware.CheckWhiteSpace(email)
		fmt.Println(whiteSpace)
		if !whiteSpace {
			fmt.Println("check", check)
			check = middleware.CheckEmail(email)
		} else {
			middleware.ErrorHandler(c, 400, "Email should not contain whitespace")
		}
	}
	return check
}

func updateToDB(username string, newPassword string, newEmail string, c *gin.Context) {
	//update statement
	//execute
	_, err := middleware.UpdateUserToDb(newPassword, newEmail, username)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"code": 200, "message": "Update success!"})
}
