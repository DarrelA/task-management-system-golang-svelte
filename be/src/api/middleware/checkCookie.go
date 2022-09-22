package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckCookie(c *gin.Context) {
	// Request for cookie
	cookie, err := c.Cookie("jwt-cookie")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized User, cookie not found"})
		return
	}

	// ParseWithClaims verify the jwt in the cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(LoadENV("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"code": http.StatusUnauthorized, "message": "Unauthorized User"})
		return
	}

	// Ignore this for now
	token.Claims.Valid()

	c.Next()
}
