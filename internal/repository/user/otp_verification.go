package user_repository

import (
	"errors"
	"log"
	// "log"
	"net/http"
	"time"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/utils"
	"gorm.io/gorm"
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
	return &pv, nil
}

// func GetVerification(db *gorm.DB, email, phone string) (*models.OtpVerification, error) {
// 	var pv models.OtpVerification
// 	err := db.Where("email = ? OR phone = ?", email, phone).First(&pv).Error
// 	if err != nil {
// 	    if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, nil 
// 		}
// 		return nil, err
// 	}
// 	return &pv, nil
// }

func UpdateVerification(db *gorm.DB, verification models.OtpVerification) error {
	verification.UpdatedAt = time.Now()
	if err := db.Save(verification).Error; err != nil {
		return err
	}
	return nil
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
		return tx.Save(&pv).Error
	})
}


func OTPCheck(db *gorm.DB, identifier string, otp string)(string, any, int, error){
	isPhone := identifier == "phone"
	verification, err := func() (*models.OtpVerification, error) {
		if isPhone {
            return GetVerification(db, identifier)
		}
        return GetVerification(db, identifier)
	}()

	log.Println("Verification", verification)

	if err != nil || verification == nil {
		return "Invalid verification", nil, http.StatusUnavailableForLegalReasons, errors.New("Invalid verification")
	}

	if !verification.IsActive {
		return "OTP deactivated", nil, http.StatusBadRequest, errors.New("Invalid OTP")
	}

	// =================== Check FOR EXPIRED OTP ========================

	if IsOTPExpired(verification.UpdatedAt, 10*time.Minute) {
		return "OTP expired", nil, http.StatusBadRequest, errors.New("OTP expired")
	}

	// ====================== GET OTP VERIFICATION =====================
	if verification.Token != otp {
		return "OTP does not match", nil, http.StatusBadRequest, errors.New("OTP does not match")
	}

	// ===================== DEACTIVATE OTP ========================
	verification.IsActive = false
	UpdateVerification(db, *verification)
	return "", verification, 200, nil
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

