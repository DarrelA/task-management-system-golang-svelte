package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Go struct in the form of JSON
type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Usergroup string `json:"usergroup"`
	Status    string `json:"status"`
}

type accounts struct {
	users User
}

var db *sql.DB
var err error

func main() {

	connectionToMySQL()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{code: 200}`)
	})

	http.HandleFunc("/admin-update-user", adminUpdateUserController)

	fmt.Printf("Starting server at port 4000\n")
	err := http.ListenAndServe(":4000", nil)
	checkError(err)

}

func adminUpdateUserController(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// returns a slice of bytes
	body, _ := ioutil.ReadAll(req.Body)

	keyVal := make(map[string]string)
	// func Unmarshal(data []byte, v interface{}) error
	json.Unmarshal(body, &keyVal)

	// Set Headers for response, server informs client that JSON data is being sent.
	w.Header().Set("Content-Type", "application/json")

	username := strings.TrimSpace(keyVal["username"])
	password := keyVal["password"]
	email := strings.TrimSpace(keyVal["email"])
	// usergroup := keyVal["usergroup"]
	// status := keyVal["status"]

	// Check if username exists
	if username == "" {
		http.Error(w, "Please enter a username.", 500)
	} else {
		rows, err := db.Query(`SELECT * FROM accounts WHERE username = ?;`, username)
		checkError(err)

		if rows.Next() {
			if password != "" {
				// _, err := db.Query(`UPDATE accounts SET password = ? WHERE username = ?`, password, username)
				// checkError(err)
				// io.WriteString(w, "Updated Password!")
				if email != "" {
					rows, err := db.Query(`SELECT * FROM accounts WHERE email = ?;`, email)
					checkError(err)
					if rows.Next() {
						http.Error(w, "This email already exist. Please try again.", 404)
					} else {
						io.WriteString(w, "Don't have this email, can update!")
					}
				}
			}
		} else {
			http.Error(w, "This username does not exist, please try again", 404)
		}
	}

	//////

	// u1 := "lowjiewei"
	// jsonData := map[string]string{
	// 	"username": u1,
	// }

	// json.NewEncoder(w).Encode(&keyVal)
	// fmt.Println(keyVal["username"])

	// stmt, err := db.Prepare(`UPDATE accounts SET password="Password123456" WHERE username="admin";`)
	// checkError(err)
	// defer stmt.Close()

	// r, err := stmt.Exec()
	// checkError(err)

	// n, err := r.RowsAffected()
	// checkError(err)

	// fmt.Fprintln(w, "UPDATED RECORD", n, "SUCCESSFULLY UPDATED")

	// defer db.Close()

	// Testing
	// if username != "" {
	// 	json.NewEncoder(w).Encode(&keyVal)
	// }

	// func NewEncoder(w io.Writer) *Encoder
	// func (enc *Encoder) Encode(v any) error
	// Conversion of Go values to JSON

}

func connectionToMySQL() {
	// db, err := sql.Open(driver, dataSourceName)
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/C3_database")
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Connected to MySQL Database!")

}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
