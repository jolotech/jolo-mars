package email

func (s *ender) Verification() error {
	body := verificationTemplate(s.User.Name, s.OTP)
	return sendMail(s.User.Email, "Email Verification", body)
}

func (s *Sender) Welcome() error {
	body := welcomeTemplate(s.User.Name)
	return sendMail(s.User.Email, "Welcome 🎉", body)
}

func (s *Sender) ForgetPassword() error {
	body := forgetPasswordTemplate(s.User.Name, s.Token)
	return sendMail(s.User.Email, "Forgot Password", body)
}

func (s *Sender) ResetPassword() error {
	body := resetPasswordTemplate(s.User.Name)
	return sendMail(s.User.Email, "Password Reset Successful", body)
}
