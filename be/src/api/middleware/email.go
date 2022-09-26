package middleware

import (
	"net/mail"
	"fmt"
)

// email validation
func CheckEmail(email string) bool {

	var (
		validEmail = false
	)

	// Parse email format before query
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println(err)
		validEmail = false
	} else {
		validEmail = true
	}
	
	return validEmail
}
