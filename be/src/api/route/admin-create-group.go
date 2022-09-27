package route

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// middleware package
	"backend/api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

type Groupnames struct {
	// json tag to de-serialize json body
	LoggedInUser string `json:"loggedInUser"`
	Name         string `json:"user_group"`
}

func AdminCreateGroup(context *gin.Context) {
	var newGroup Groupnames

	// call BindJSON to bind the received JSON to newGroup
	if err := context.BindJSON(&newGroup); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(context, 400, "Bad Request")
		return
	}

	// Check user group
	checkGroup := middleware.CheckGroup(context.GetString("username"), "Admin")
	if !checkGroup {
		middleware.ErrorHandler(context, 400, "Unauthorized actions")
		return
	}

	// check if groupname field has whitespace
	whiteSpace := middleware.CheckWhiteSpace(newGroup.Name)
	if whiteSpace {
		middleware.ErrorHandler(context, 400, "Groupname should not contain whitespace")
		return
	}

	// check if groupname field is empty
	minLength := middleware.CheckLength(newGroup.Name)
	if minLength {
		middleware.ErrorHandler(context, 400, "Groupname should not be empty")
		return
	}

	// check for existing groupname before creating
	// return first result (single row result)
	result := middleware.SelectUserGroupFromGroupnamesByUserGroup(newGroup.Name)

	// Scan: scanning and reading input (texts given in standard input)
	switch err := result.Scan(&newGroup.Name); err {

	// New Group
	case sql.ErrNoRows:
		// insert new group
		_, err := middleware.InsertGroupnames(newGroup.Name)

		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(context, 400, "Unable to create new group")
			return
		}

		context.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "New group has created successfully"})
		return

	// Existing groupname
	case nil:
		middleware.ErrorHandler(context, 400, "Existing Groupname")
		return

	// Invalid Field
	default:
		middleware.ErrorHandler(context, 400, "Invalid Field")
		return
	}
}
