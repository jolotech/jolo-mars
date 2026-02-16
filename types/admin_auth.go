package types

type AdminLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AdminChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=10"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=10"`
	Email           string `json:"email,omitempty"`
}

type AdminLoginResponse struct {
	AccessToken            string `json:"access_token,omitempty"`
	Requires2FA bool   `json:"requires_2fa"`
	Requires2FAMessage string `json:"requires_2fa_message,omitempty"`
	TwoFAToken  string `json:"two_fa_token,omitempty"`
	PasswordChangeRequired bool   `json:"password_change_required"`
	SetupToken             string `json:"setup_token,omitempty"`
	Admin                  any    `json:"admin,omitempty"`
}

type AdminTwoFASetupResponse struct {
	OtpAuthURL string `json:"otpauth_url"`
}

type AdminTwoFAConfirmRequest struct {
	Code string `json:"code" binding:"required,len=6,numeric"`
}

type AdminLogin2FARequest struct {
	TwoFAToken string `json:"two_fa_token" binding:"required"`
	Code       string `json:"code" binding:"required,len=6,numeric"`
}