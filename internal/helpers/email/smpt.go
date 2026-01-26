package email

import (
	"net/smtp"
)

const (
	SMTPHost = "smtp.zoho.com"
	SMTPPort = "587"

	SMTPUser = "no-reply@yourdomain.com"
	SMTPPass = "YOUR_ZOHO_APP_PASSWORD"
)

func sendMail(to, subject, body string) error {
	auth := smtp.PlainAuth(
		"",
		SMTPUser,
		SMTPPass,
		SMTPHost,
	)

	msg := []byte(
		"From: " + SMTPUser + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			body,
	)

	return smtp.SendMail(
		SMTPHost+":"+SMTPPort,
		auth,
		SMTPUser,
		[]string{to},
		msg,
	)
}
