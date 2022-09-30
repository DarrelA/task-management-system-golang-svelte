package middleware

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Subject string
	Body    string
}

// include taskname, sender email, recipient email as param
func SendMail(c *gin.Context) {
	username := LoadENV("SMTP_USERNAME")
	password := LoadENV("SMTP_PASSWORD")
	host := LoadENV("SMTP_HOST")

	recipient := []string{"project_lead@tms.com", "project_lead2@tms.com"}
	sender := "team_member@tms.com"
	cc := []string{}

	auth := smtp.PlainAuth("", username, password, host)

	subject := "New task have been completed!"
	body := "<h3>Task has been completed by team member.</h3>\r\n" +
		"Review Now <task name>!\r\n"

	mail := Mail{
		Sender:  sender,
		To:      recipient,
		Cc:      cc,
		Subject: subject,
		Body:    body,
	}

	msg := BuildMessage(mail)

	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, sender, recipient, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Email sent successfully")
}

// Compose messsage template
func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)

	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))

	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
