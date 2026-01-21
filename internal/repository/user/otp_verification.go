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
