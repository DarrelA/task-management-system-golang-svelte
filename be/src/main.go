package main

import (
	"backend/api/middleware"
	"backend/api/route"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var db *sql.DB

func main() {
	// route.Login()

	db := middleware.ConnectionToDatabase()

	defer db.Close()

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/login", route.Login)
	router.POST("/logout", route.Logout)

	// router.POST("/admin-update-user", middleware.CheckCookie, route.AdminUpdateUser)
	router.POST("/admin-update-user", route.AdminUpdateUser)
	router.POST("/admin-create-user", route.AdminCreateUser)
	router.POST("/admin-create-group", route.AdminCreateGroup)
	router.GET("/get-users", route.GetUsers)
	router.GET("/get-user-groups", middleware.GetUserGroup)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
