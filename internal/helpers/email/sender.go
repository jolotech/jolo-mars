package email

import (
	// "fmt"
	// "net/smtp"

	"github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/types"
)



type EmailSender struct {
	OTP   string
	Token string
	User  *models.User
	ToEmail string
	ToName  string
}


func SendEmail(value interface{}, user *models.User) *EmailSender {
	s := &EmailSender{User: user}

	if v, ok := value.(string); ok {
		s.OTP = v
		s.Token = v
	}

	return s
}


func SendAdminEmail(toEmail, toName string) *EmailSender {
	return &EmailSender{
		ToEmail: toEmail,
		ToName:  toName,
	}
}