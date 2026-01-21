package user_repository

import (
	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
)

func GetPhoneVerification(db *gorm.DB, phone string) (*models.PhoneVerification, error) {
	var pv models.PhoneVerification

	err := db.Where("phone = ?", phone).First(&pv).Error
	if err != nil {
		return nil, err
	}

	return &pv, nil
}


func UpsertPhoneOTP(db *gorm.DB, phone, otp string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var pv models.PhoneVerification

		err := tx.Where("phone = ?", phone).First(&pv).Error
		if err != nil {
			// insert
			return tx.Create(&models.PhoneVerification{
				Phone:       phone,
				Token:       otp,
				OtpHitCount: 0,
			}).Error
		}

		// update
		pv.Token = otp
		pv.OtpHitCount = 0
		return tx.Save(&pv).Error
	})
}
func DeletePhoneVerification(db *gorm.DB, phone string) error {
	return db.Where("phone = ?", phone).Delete(&models.PhoneVerification{}).Error
}