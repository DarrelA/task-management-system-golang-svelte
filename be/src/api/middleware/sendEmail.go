package middleware

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {
	username := LoadENV("SMTP_USERNAME")
	password := LoadENV("SMTP_PASSWORD")
	host := LoadENV("SMTP_HOST")

	auth := smtp.PlainAuth("", username, password, host)

	to := []string{"project_lead@tms.com"}
	from := "team_member@tms.com"
	msg := []byte("To: project_lead@tms.com\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here's the space for our great sales pitch\r\n")

	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Email sent, %s", to)
}
