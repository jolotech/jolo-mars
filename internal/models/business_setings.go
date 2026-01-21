package models

import (
	"time"
	// "gorm.io/gorm"
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