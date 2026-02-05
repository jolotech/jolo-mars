package types


// import "github.com/jolotech/jolo-mars/internal/models"

// type EmailSender struct {
// 	OTP   string
// 	Token string
// 	User  models.User
// }

type Sender struct {
	User *struct {
		Name  string
		Email string
	}
	OTP   string
	Token string
}