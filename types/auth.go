package types


type RegisterRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required,min=14"`
	Password  string `json:"password" binding:"required,min=8"`
	RefCode   string `json:"ref_code"`
	OtpOption string `json:"otp_option"`
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

type VerifyOTPRequest struct {
	OTP              string `json:"otp" binding:"required"`
	VerificationMethod string `json:"verification_method" binding:"required,oneof=phone email"`
	Phone            string `json:"phone,omitempty"`
	Email            string `json:"email,omitempty"`
	GuestID          *string `json:"guest_id,omitempty"`
}

type ResendOTPRequest struct {
	VerificationMethod string `json:"verification_method" binding:"required,oneof=phone email"`
	Phone            string `json:"phone,omitempty"`
	Email            string `json:"email,omitempty"`
}

type ResetPasswordSubmitRequest struct {
	ResetToken         string `json:"reset_token" binding:"required"`
	Password           string `json:"password" binding:"required,min=8"`
	ConfirmPassword    string `json:"confirm_password" binding:"required"`
	VerificationMethod string `json:"verification_method" binding:"required,oneof=phone email"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
}

type GuestRequest struct {
	FCMToken  string `json:"fcm_token" binding:"required"`
	IPAddress string `json:"ip_address,omitempty"`
}

type GuestResponse struct {
	GuestID uint `json:"guest_id"`
}