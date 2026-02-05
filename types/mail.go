package types


// import "github.com/jolotech/jolo-mars/internal/models"

// type EmailSender struct {
// 	OTP   string
// 	Token string
// 	User  models.User
// }
type EmailUser struct {
	Name  string
	Email string
}

type Sender struct {
	User EmailUser
	OTP   string
	Token string
}