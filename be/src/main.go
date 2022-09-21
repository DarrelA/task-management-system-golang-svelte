package main

import (
	"backend/api/middleware"
	"backend/api/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

// var db *sql.DB

func main() {
	// route.Login()

	db := middleware.ConnectionToDatabase()

	defer db.Close()

	router := gin.Default()

	router.POST("/admin-update-user", route.AdminUpdateUserController)
	router.POST("/admin-create-user", route.AdminCreateUser)
	router.GET("/get-users", route.GetUsers)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
