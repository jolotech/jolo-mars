package email

import (
	// "github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/types"
)


func (s *EmailSender) Verification() error {
	body, err := renderTemplate("verification.html", map[string]any{
		"Name": s.User.FName,
		"OTP":  s.OTP,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.FName, "Email Verification", body)
}

func (s *EmailSender) Welcome() error {
	body, err := renderTemplate("welcome.html", map[string]any{
		"Name": s.User.FName,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.FName, "Welcome 🎉", body)
}

func (s *EmailSender) ForgetPassword() error {
	body, err := renderTemplate("forget_password.html", map[string]any{
		"Name":     s.User.FName,
		"ResetURL": "https://shop.jolojolo.com/reset-password?token=" + s.Token,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.FName, "Reset Your Password", body)
}

func (s *EmailSender) ResetPassword() error {
	body, err := renderTemplate("reset_password.html", map[string]any{
		"Name": s.User.FName,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.FName, "Password Reset Successful", body)
}
