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
	router.GET("/logout", route.Logout)

	// router.GET("/check-group", middleware.CheckGroup(username, groupname))
	router.POST("/add-user-to-group", middleware.CheckCookie, route.AddUserToGroup)
	router.POST("/admin-update-user", middleware.CheckCookie, route.AdminUpdateUser)
	router.POST("/admin-create-user", middleware.CheckCookie, route.AdminCreateUser)
	router.POST("/admin-create-group", middleware.CheckCookie, route.AdminCreateGroup)
	router.GET("/get-users", middleware.CheckCookie, route.GetUsers)
	router.GET("/get-user-groups", middleware.CheckCookie, route.GetUserGroup)

	// router.GET("/get-users-in-group", route.GetUsersInGroup)
	router.GET("/get-users-in-group", middleware.CheckCookie, route.GetUsersInGroup)

	router.POST("/update-user", route.UpdateUser)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
