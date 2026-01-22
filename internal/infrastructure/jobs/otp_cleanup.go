package jobs

import (
	"time"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/utils"
)

func CleanupExpiredOTPs(db *gorm.DB) error {
	expiryTime := time.Now().Add(-utils.OTPExpiryMinutes * time.Minute)

	if err := db.Where("updated_at < ?", expiryTime).
		Delete(&models.OtpVerification{}).Error; err != nil {
		return err
	}
	return nil
}