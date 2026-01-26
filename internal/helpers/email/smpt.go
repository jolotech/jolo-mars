package email

import (
	"fmt"
	"net/smtp"
	"github.com/jolotech/jolo-mars/config"
)

func sendMail(to, subject, body string) error {

	cfg := config.LoadConfig()

	auth := smtp.PlainAuth(
		"",
		cfg.SMTPUser,
		cfg.SMTPPass,
		cfg.SMTPHost,
	)

	fmt.Println("Sending email to:", to, "SMTP")


	msg := []byte(
		"From: " + cfg.SMTPUser + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			body,
	)


	if err := smtp.SendMail(
		cfg.SMTPHost+":"+cfg.SMTPPort,
		auth,
		cfg.SMTPUser,
		[]string{to},
		msg,
	); err != nil {
		return fmt.Errorf("smtp send failed: %w", err)
	}

	return nil
}
