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
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": "Invalid Credentials"})
		return
	}

	// ParseWithClaims verify the jwt in the cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(LoadENV("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": "Invalid Credentials"})
		return
	}

	// set `claims` variable to be a type of .RegisteredClaims`
	// so we can access `claims.Issuer`
	claims := token.Claims.(*jwt.RegisteredClaims)

	// basically issuer is our user id
	// Refer to middleware.Login
	c.Set("username", claims.Issuer)
	c.Next()
}
