package middleware

import (
	"database/sql"
	"log"
)

func GetUserGroup() {
	db, err := sql.Open("mysql", "root:admin123@/c3_database")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// SELECT * FROM usergroup
	result, err := db.Query("SELECT * FROM groupnames")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()
	for result.Next() {
		var usergroup string

		if err := result.Scan(&usergroup); err != nil {
			log.Fatal(err)
		}
		log.Println(usergroup)
	}

}
