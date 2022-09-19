package middleware

import "net/mail"

// email validation
func CheckEmail(email string) bool {

	// Optional email
	if len(email) == 0 {
		return true
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
