package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Function to load env file values based on key param
func LoadENV(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Some other error occurred", err)
	}
	return os.Getenv(key)
}
