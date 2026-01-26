package types


import "github.com/jolotech/jolo-mars/internal/models"

type EmailSender struct {
	OTP   string
	Token string
	User  models.User
}
