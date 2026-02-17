package email

import (
	"fmt"
	// "github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/types"
)

func (s *EmailSender) AdminBootstrapCredentials(appName, role, tempPassword, loginURL, surpportEmail string) error {
	name := s.ToName
	if name == "" {
		name = "Admin"
	}

	body, err := renderTemplate("create_super_admin_credentials.html", map[string]any{
		"Name":         name,
		"Email":        s.ToEmail,
		"TempPassword": tempPassword,
		"Role":         role,
		"AppName":      appName,
		"LoginURL":     loginURL,
		"SupportEmail": surpportEmail,
	})
	if err != nil {
		return err
	}

	subject := fmt.Sprintf("%s Admin Access Details", appName)
	return sendMail(s.ToEmail, subject, body)
}


func (s *EmailSender) Verification() error {
	body, err := renderTemplate("verification.html", map[string]any{
		"Name": s.User.FName,
		"OTP":  s.OTP,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Email Verification", body)
}

func (s *EmailSender) Welcome() error {
	body, err := renderTemplate("welcome.html", map[string]any{
		"Name": s.User.FName,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Welcome To Jolo Jolo 🎉", body)
}

func (s *EmailSender) ForgetPassword() error {
	body, err := renderTemplate("forget_password.html", map[string]any{
		"Name":     s.User.FName,
		"ResetURL": "https://shop.jolojolo.com/reset-password?token=" + s.Token,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Reset Your Password", body)
}

func (s *EmailSender) ResetPassword() error {
	body, err := renderTemplate("reset_password.html", map[string]any{
		"Name": s.User.FName,
	})
	if err != nil {
		return err
	}

	return sendMail(s.User.Email, "Password Reset Successful", body)
}
