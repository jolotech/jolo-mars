package types


type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	RefCode  string `json:"ref_code"`
}

type LoginSettings struct {
	ManualLogin       bool `json:"manual_login_status"`
	OtpLogin          bool `json:"otp_login_status"`
	SocialLogin       bool `json:"social_login_status"`
	GoogleLogin       bool `json:"google_login_status"`
	FacebookLogin     bool `json:"facebook_login_status"`
	AppleLogin        bool `json:"apple_login_status"`
	EmailVerification bool `json:"email_verification_status"`
	PhoneVerification bool `json:"phone_verification_status"`
}