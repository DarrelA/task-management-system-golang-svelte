package middleware

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"backend/api/models"

	"github.com/gin-gonic/gin"
)

// include taskname, sender email, recipient email as param
func SendMail(c *gin.Context, recipient []string, senderEmail string, senderName string, taskName string) {
	username := LoadENV("SMTP_USERNAME")
	password := LoadENV("SMTP_PASSWORD")
	host := LoadENV("SMTP_HOST")

	// recipient := []string{"project_lead@tms.com", "project_lead2@tms.com"}
	// sender := "team_member@tms.com"
	cc := []string{}

	auth := smtp.PlainAuth("", username, password, host)

	subject := "New task have been completed!"
	// body := fmt.Sprintf("<h3 style='font-family': Montserrat;>Task has been completed by %s.</h3>\r\n"+
	// 	"<p style='font-family': Montserrat;>Review task: %s in TMS now!</p>\r\n", sender, taskName)

	body := fmt.Sprintf("<h3>Task has been completed by %s.</h3>\r\n"+
		"<p>Review task: %s in TMS now!</p>\r\n", senderName, taskName)

	mail := models.Email{
		Sender:  senderEmail,
		To:      recipient,
		Cc:      cc,
		Subject: subject,
		Body:    body,
	}

	message := BuildMessage(mail)

	// host:port, auth, from, to, []byte
	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, senderEmail, recipient, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Email sent successfully")
}

// Compose messsage template
func BuildMessage(mail models.Email) string {
	message := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	message += fmt.Sprintf("From: %s\r\n", mail.Sender)

	message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))

	if len(mail.Cc) > 0 {
		message += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return message
}
