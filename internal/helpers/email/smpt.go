package email

import (
	"net/smtp"
	"github.com/jolotech/jolo-mars/config"
)

// const (
// 	SMTPHost = "smtp.zoho.com"
// 	SMTPPort = "587"

// 	SMTPUser = "no-reply@yourdomain.com"
// 	SMTPPass = "YOUR_ZOHO_APP_PASSWORD"
// )

func sendMail(to, subject, body string) error {

	cfg := config.LoadConfig()

	auth := smtp.PlainAuth(
		"",
		cfg.SMTPUser,
		cfg.SMTPPass,
		cfg.SMTPHost,
	)

	msg := []byte(
		"From: " + cfg.SMTPUser + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			body,
	)

	return smtp.SendMail(
		cfg.SMTPHost+":"+cfg.SMTPPort,
		auth,
		cfg.SMTPUser,
		[]string{to},
		msg,
	)
}
