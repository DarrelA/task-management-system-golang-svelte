package middleware

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// password validation
func CheckPassword(password string) bool {
	var (
		hasMinLength = false
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		hasSpecial   = false
	)

	// check length of password
	if len(password) >= 8 && len(password) <= 10 {
		hasMinLength = true
	}

	// check contains special characters
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSpecial = true
		}
	}
	return hasMinLength && hasUpper && hasLower && hasNumber && hasSpecial
}

// password hashing
func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
