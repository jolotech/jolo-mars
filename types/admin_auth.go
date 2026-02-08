package types

type AdminLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AdminChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
	EMAIL           string `json:"email,omitempty"`
}

type AdminLoginResponse struct {
	AccessToken            string `json:"access_token,omitempty"`
	PasswordChangeRequired bool   `json:"password_change_required"`
	SetupToken             string `json:"setup_token,omitempty"` // used only for forced password change
	Admin                  any    `json:"admin,omitempty"`
}