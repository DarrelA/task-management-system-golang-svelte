package main

import (
	"backend/api/middleware"
	"backend/api/route"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var db *sql.DB

func main() {
	// route.Login()

	db := middleware.ConnectionToDatabase()

	defer db.Close()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // For authentication
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.POST("/login", route.Login)
	router.POST("/logout", route.Logout)

	// Dont use checkgroup like this. Refer to admin-create-user line 46-50
	// router.GET("/check-group", middleware.CheckGroup)
	router.POST("/user-to-group", route.AddUserToGroup)
	router.POST("/admin-update-user", middleware.CheckCookie, route.AdminUpdateUser)
	router.POST("/admin-create-user", middleware.CheckCookie, route.AdminCreateUser)
	router.POST("/admin-create-group", route.AdminCreateGroup)
	router.GET("/get-users", route.GetUsers)
	router.GET("/get-user-groups", route.GetUserGroup)
	router.POST("/update-user", route.UpdateUserFunction)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
