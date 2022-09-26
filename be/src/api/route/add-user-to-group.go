package route

import (
	"backend/api/middleware"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CompositeKey struct {
	Username  string `json:"username"`
	Groupname string `json:"groupname"`
}

func AddUserToGroup(c *gin.Context) {
	// Bind request JSON
	var newComposite CompositeKey
	if err := c.BindJSON(&newComposite); err != nil {
		fmt.Println(err)
		middleware.ErrorHandler(c, 400, "Bad Request")
		return
	}

	// Fetch exisiting usernames
	// rows, err := db.Query("SELECT username FROM accounts")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// var username string
	// var data []string

	// for rows.Next() {
	// 	err = rows.Scan(&username)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Append to slice
	// 	data = append(data, username)
	// }
	// // Send data to frontend
	// c.JSON(200, data)

	// Check if usergroup exist

	var user_group string
	fmt.Println("1", newComposite.Groupname)
	getGroupname := "SELECT user_group FROM groupnames WHERE user_group = ?"
	group := db.QueryRow(getGroupname, newComposite.Groupname)

	switch err := group.Scan(&user_group); err {

	// New group name
	case sql.ErrNoRows:
		// INSERT user_group into groupnames table
		_, err := db.Exec("INSERT INTO groupnames (user_group) VALUES (?)", newComposite.Groupname)
		if err != nil {
			fmt.Println(err)
			middleware.ErrorHandler(c, 400, "Invalid field")
			return
		}
	}

	// Create a composite key for usergroup table
	_, err := db.Exec("INSERT INTO usergroup (username, user_group) VALUES (?, ?)", newComposite.Username, newComposite.Groupname)
	if err != nil {
		fmt.Println("Here")
		fmt.Println(err)
		return
	}

	fmt.Printf("Added new composite key %v to %v \n", newComposite.Username, newComposite.Groupname)
	response := fmt.Sprintf("%s added to group: %s", newComposite.Username, newComposite.Groupname)

	// Fetch user's EXISTING groups and update
	var existingGroup string
	getUser := "SELECT user_group FROM accounts WHERE username = ?"
	user := db.QueryRow(getUser, newComposite.Username)
	err = user.Scan(&existingGroup)
	if err != nil {
		panic(err)
	}
	newGroup := fmt.Sprintf("%s, %s", existingGroup, newComposite.Groupname)

	// UPDATE statement
	_, err = db.Exec("UPDATE accounts SET user_group = ? WHERE username = ?", newGroup, newComposite.Username)
	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully updated accounts table")

	c.JSON(201, response)

}
