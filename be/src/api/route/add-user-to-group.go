package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type CompositeKey struct {
	Username     string   `json:"username"`
	Groupname    []string `json:"groupname"`
}

var existingGroup string

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
	fmt.Println(checkGroup)
	if !checkGroup {
		middleware.ErrorHandler(c, 400, "Unauthorized actions")
		return
	}

	for _, group := range newComposite.Groupname {
		// LOOP to validate group name
		var user_group string

		// "SELECT user_group FROM groupnames WHERE user_group = ?;"
		result := middleware.SelectGroupnamesbyUserGroup(group)
		// fmt.Println(newComposite.Username, group)

		switch err := result.Scan(&user_group); err {

		// New group name
		case sql.ErrNoRows:
			// INSERT user_group into groupnames table
			_, err := middleware.InsertGroupnames(group)
			if err != nil {
				middleware.ErrorHandler(c, 400, "Invalid field")
				return
			}

			_, err = middleware.InsertUserGroup(newComposite.Username, group)
			// fmt.Println("New Composite key", newComposite.Username, group)
			if err != nil {
				fmt.Println(err)
				return
			}
		// Existing group name
		case nil:
			// Check if composite key exist
			// "SELECT username, user_group FROM usergroup WHERE username = ? AND user_group = ?"
			key := middleware.SelectCompositeKey(newComposite.Username, group)
			switch err := key.Scan(&newComposite.Username, group); err {
			case sql.ErrNoRows:
				// Create a composite key for usergroup table
				_, err := middleware.InsertUserGroup(newComposite.Username, group)
				// fmt.Println("New Composite key", newComposite.Username, group)
				if err != nil {
					fmt.Println(err)
					return
				}  
			}
		}

		// Fetch user's EXISTING groups and update
		
		// SELECT user_group FROM accounts WHERE username = ?
		result = middleware.SelectUserFromUserGroupByUsername(newComposite.Username)
		err := result.Scan(&existingGroup)
	
		if err != nil {
			panic(err)
		}		
	}

	groupstr := strings.Join(newComposite.Groupname, ",")
	fmt.Println(groupstr)

	existingGroup = existingGroup + "," + groupstr

	groupSlice := strings.Split(existingGroup, ",")
	fmt.Println(groupSlice)

	keys := make(map[string]bool )

	// composite groupname slice
	list := []string{}

	for _, entry := range groupSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	groupstr = strings.Join(list, ",")
	fmt.Println(groupstr)

	_, err := middleware.UpdateAccountsSetUsernameByUsergroup(groupstr, newComposite.Username)
	if err != nil {
		panic(err)
	}

	c.JSON(201, gin.H{
		"code":    201,
		"message": "Added user to group",
	})
}
