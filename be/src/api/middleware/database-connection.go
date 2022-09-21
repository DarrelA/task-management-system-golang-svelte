package middleware

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectionToDatabase() *sql.DB {
	// db, err := sql.Open(driver, dataSourceName)
	dataSourceName := fmt.Sprintf("%s:%s@/%s", LoadENV("SERVER_USER"), LoadENV("SERVER_PASSWORD"), LoadENV("SERVER_DB"))
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}
	fmt.Println("Connected to MySQL Database!")

	return db
}
