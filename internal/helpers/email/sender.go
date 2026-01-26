package email

import (
	// "fmt"
	// "net/smtp"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/types"
)

// type EmailSender struct {
// 	User  *models.User
// 	OTP   string
// 	Token string
// }

// func SendEmail(value interface{}, user *models.User) *types.EmailSender {
// 	sender := &types.EmailSender{}
// 	if user != nil {
// 		sender.User = *user
// 	}

// 	// decide what was passed
// 	switch v := value.(type) {
// 	case string:
// 		sender.OTP = v
// 		sender.Token = v
// 	}

// 	return sender
// }


func SendEmail(value interface{}, user *models.User) *types.EmailSender {

	sender := &types.EmailSender{}
	if user != nil {
		sender.User = *user
	}

	if v, ok := value.(string); ok {
		sender.OTP = v
		sender.Token = v
	}

	return sender
}
