package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type CompositeKey struct {
	LoggedInUser string   `json:"loggedInUser"`
	Username     string   `json:"username"`
	Groupname    []string `json:"groupname"`
}

func AddUserToGroup(c *gin.Context) {
	// Bind request JSON
	var newComposite CompositeKey
	if err := c.BindJSON(&newComposite); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}

	// Check user group
	checkGroup := middleware.CheckGroup(c.GetString("username"), "Admin")
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	for _, group := range newComposite.Groupname {
		// LOOP to validate group name
		var user_group string

		result := middleware.SelectGroupnamesbyUserGroup(group)

		switch err := result.Scan(&user_group); err {

		// New group name
		case sql.ErrNoRows:
			// INSERT user_group into groupnames table
			_, err := middleware.InsertGroupnames(group)
			if err != nil {
				middleware.ErrorHandler(c, 400, "Invalid field")
				return
			}
			fmt.Println("New group name")
		}

		// EXISTING GROUP NAME
		// fmt.Println("existing group name")

		// Check if composite key exist
		key := middleware.SelectCompositeKey(newComposite.Username, group)
		switch err := key.Scan(&newComposite.Username, group); err {
		case sql.ErrNoRows:
			// Create a composite key for usergroup table
			_, err := middleware.InsertUserGroup(newComposite.Username, group)
			fmt.Println("New Composite key")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Fetch user's EXISTING groups and update
		var existingGroup string
		result = middleware.SelectUserFromUserGroupByUsername(newComposite.Username)
		err := result.Scan(&newComposite.Username)
		if err != nil {
			panic(err)
		}
		containsGroup := strings.Contains(existingGroup, group)
		if !containsGroup {
			newGroup := fmt.Sprintf("%s, %s", existingGroup, group)
			_, err := middleware.UpdateAccountsSetUsernameByUsergroup(newGroup, newComposite.Username)
			if err != nil {
				panic(err)
			}
		}
	}

	c.JSON(201, gin.H{
		"code":    201,
		"message": "Added user to group",
	})
}
