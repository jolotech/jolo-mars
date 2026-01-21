package admin_repository

import (
	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/types"
)

type Main struct {
	db *gorm.DB
}	


func NewAdminMainRepository(db *gorm.DB) *Main {
	return &Main{db: db}
}


func (r *Main) GetBusinessSetting(key string) interface{} {
	var setting models.BusinessSetting
	result := r.db.First(&setting)
	if result.Error != nil {
		return nil
	}
	switch key {
	case "ref_earning_status":
		return setting.RefEarningStatus
	case "registration_bonus_status":
		return setting.RegistrationBonusStatus
	case "registration_bonus_amount":
		return setting.RegistrationBonusAmount
	case "service_charge_percent":
		return setting.ServiceChargePercent
	case "firebase_otp_verification":
		return setting.FirebaseOTPVerification
	default:
		return nil
	}
}


func (r *Main) GetLoginSettings() *types.LoginSettings {
	var bs models.BusinessSetting

	if err := r.db.First(&bs).Error; err != nil {
		return nil
	}

	return &types.LoginSettings{
		ManualLogin:       bs.ManualLoginStatus,
		OtpLogin:          bs.OtpLoginStatus,
		SocialLogin:       bs.SocialLoginStatus,
		GoogleLogin:       bs.GoogleLoginStatus,
		FacebookLogin:     bs.FacebookLoginStatus,
		AppleLogin:        bs.AppleLoginStatus,
		EmailVerification: bs.EmailVerificationStatus,
		PhoneVerification: bs.PhoneVerificationStatus,
	}
}