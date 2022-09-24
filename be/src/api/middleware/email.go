package middleware

import (
	"database/sql"
	"net/mail"
)

// email validation
func CheckEmail(email string) bool {

	var (
		validEmail = false
	)

	// Check if email exist
	checkEmail := "SELECT email FROM accounts WHERE email = ?"
	result := db.QueryRow(checkEmail, email)

	switch err := result.Scan(&email); err {

	// Email dont exist
	case sql.ErrNoRows:

		_, err := mail.ParseAddress(email)
		if err != nil {
			validEmail = false
		}
		validEmail = true

	case nil:
		validEmail = false
	}

	return validEmail
}
