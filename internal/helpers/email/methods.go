package email

import (
	// "github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/types"
)

type Sender struct {
	User *struct {
		Name  string
		Email string
	}
	OTP   string
	Token string
}

func (s *Sender) Verification() error {
	body, err := renderTemplate("verification.html", map[string]any{
		"Name": s.User.Name,
		"OTP":  s.OTP,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Email Verification", body)
}

func (s *Sender) Welcome() error {
	body, err := renderTemplate("welcome.html", map[string]any{
		"Name": s.User.Name,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Welcome 🎉", body)
}

func (s *Sender) ForgetPassword() error {
	body, err := renderTemplate("forget_password.html", map[string]any{
		"Name":     s.User.Name,
		"ResetURL": "https://shop.jolojolo.com/reset-password?token=" + s.Token,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Reset Your Password", body)
}

func (s *Sender) ResetPassword() error {
	body, err := renderTemplate("reset_password.html", map[string]any{
		"Name": s.User.Name,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Password Reset Successful", body)
}
