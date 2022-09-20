package middleware

import (
	"unicode"
)

func CheckWhiteSpace(value string) bool {
	var (
		hasWhiteSpace = false
	)
	// s := strings.TrimSpace(value)
	for _, char := range value {
		switch {
		case unicode.IsSpace(char) :
			hasWhiteSpace = true
		}
	} 
	return hasWhiteSpace
	
}

func CheckLength(value string) bool {
	var (
		hasMinLength = false
	)

	if len(value) == 0 {
		hasMinLength = true
		
	}
	return hasMinLength
}