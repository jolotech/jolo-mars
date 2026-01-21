package models

import (
	"time"
	"gorm.io/gorm"
)

type BusinessSetting struct {
	ID                      uint       `json:"id" gorm:"primaryKey"`

	// Referral
	RefEarningStatus        bool        `json:"ref_earning_status" gorm:"default:true"`

	// Future settings (examples)
	RegistrationBonusStatus bool        `json:"registration_bonus_status" gorm:"default:false"`
	RegistrationBonusAmount float64     `json:"registration_bonus_amount" gorm:"default:0"`
	ServiceChargePercent    float64     `json:"service_charge_percent" gorm:"default:0"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
}

func GetBusinessSetting(db *gorm.DB, key string) interface{} {
	var setting BusinessSetting
	result := db.First(&setting)
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
	default:
		return nil
	}
}