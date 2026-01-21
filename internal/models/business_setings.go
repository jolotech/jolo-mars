package models

// import (
// 	"time"
// )

// type BusinessSetting struct {
// 	ID                      uint       `json:"id" gorm:"primaryKey"`
// 	RefEarningStatus        bool        `json:"ref_earning_status" gorm:"default:true"`
// 	RegistrationBonusStatus bool        `json:"registration_bonus_status" gorm:"default:false"`
// 	RegistrationBonusAmount float64     `json:"registration_bonus_amount" gorm:"default:0"`
// 	ServiceChargePercent    float64     `json:"service_charge_percent" gorm:"default:0"`
// 	CreatedAt               time.Time   `json:"created_at"`
// 	UpdatedAt               time.Time   `json:"updated_at"`
// }



import "time"

type BusinessSetting struct {
	ID uint `gorm:"primaryKey"`

	// Referral & bonuses
	RefEarningStatus        bool    `json:"ref_earning_status" gorm:"default:false"`
	RegistrationBonusStatus bool    `json:"registration_bonus_status" gorm:"default:false"`
	RegistrationBonusAmount float64 `json:"registration_bonus_amount" gorm:"default:0"`
	ServiceChargePercent    float64 `json:"service_charge_percent" gorm:"default:0"`

	// Login settings
	ManualLoginStatus       bool    `json:"manual_login_status" gorm:"default:false"`
	OtpLoginStatus          bool    `json:"otp_login_status" gorm:"default:false"`
	SocialLoginStatus       bool    `json:"social_login_status" gorm:"default:false"`
	GoogleLoginStatus       bool    `json:"google_login_status" gorm:"default:false"`
	FacebookLoginStatus     bool    `json:"facebook_login_status" gorm:"default:false"`
	AppleLoginStatus        bool    `json:"apple_login_status" gorm:"default:false"`

	// Verification
	EmailVerificationStatus bool    `json:"email_verification_status" gorm:"default:false"`
	PhoneVerificationStatus bool    `json:"phone_verification_status" gorm:"default:false"`
	FirebaseOTPVerification bool    `json:"firebase_otp_verification" gorm:"default:false"`

	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
