package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
)

// User login details: username, password
type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// After user login, generate JWT and set cookie
func Login(c *gin.Context) {
	var credentials LoginCredentials

	// Placeholder for results from db query
	var username string
	var password string

	// Decode JSON body to credentials struct
	if err := c.BindJSON(&credentials); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, http.StatusBadRequest, "Bad Request")
		return
	}

	// Fetch username and password from db
	getUser := "SELECT username, password FROM accounts WHERE username = ?"
	result := db.QueryRow(getUser, credentials.Username)

	switch err := result.Scan(&username, &password); err {
	case sql.ErrNoRows:
		fmt.Println("Invalid Credentials")
		middleware.ErrorHandler(c, http.StatusBadRequest, "Invalid Credentials")

	case nil:
		// Username found
		// Compare password hash
		correctPassword := middleware.CompareHash(credentials.Password, password)

		// Incorrect password
		if !correctPassword {
			middleware.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized User")
			return
		}

		nowTime := time.Now()
		expireTime := nowTime.Add(1 * time.Hour)

		// New token with signing method and claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer:    username, // MySQL > Accounts Table > username
			ExpiresAt: jwt.NewNumericDate(expireTime),
		})

		tokenString, err := token.SignedString([]byte(middleware.LoadENV("JWT_SECRET")))

		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Set cookie
		c.SetCookie("token", tokenString, 10, "/", "localhost", false, true)

		c.JSON(200, gin.H{
			"code":  200,
			"token": tokenString,
		})
	}

}
