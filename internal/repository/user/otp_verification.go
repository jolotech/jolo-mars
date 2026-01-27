package user_repository

import (
	"time"
	"errors"
	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/utils"
)

// func GetVerification(db *gorm.DB, value string) (*models.OtpVerification, error) {
// 	var pv models.OtpVerification

// 	err := db.Where("verification_method = ?", value).First(&pv).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pv, nil
// }



func GetVerification(db *gorm.DB, value string) (*models.OtpVerification, error) {
	var pv models.OtpVerification

	err := db.Where("verification_method = ?", value).First(&pv).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err
	}

	return &pv, nil
}



func UpsertOTP(db *gorm.DB, value, otp string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var pv models.OtpVerification

		err := tx.Where("verification_method = ?", value).First(&pv).Error
		if err != nil {
			// insert
			return tx.Create(&models.OtpVerification{
				VerificationMethod: value,
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


func IncrementOtpHit(db *gorm.DB, value string) error {
	return db.Model(&models.OtpVerification{}).
		Where("verification_method = ?", value).
		UpdateColumn("otp_hit_count", gorm.Expr("otp_hit_count + 1")).Error
}

func  IsOtpLocked(pv *models.OtpVerification) bool {
	return pv.OtpHitCount >= utils.OTPMaxHitCount
}

func IsOTPExpired(UpdatedAt time.Time, duration time.Duration) bool {
	return time.Since(UpdatedAt) > duration
}

