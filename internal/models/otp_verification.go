package models

import "time"


type OtpVerification struct {
	ID                  uint      `gorm:"primaryKey"`
	Token               string    `gorm:"size:6;not null"`
	VerificationMethod  string    `gorm:"not null"`
	IsActive            bool      `gorm:"default:true"`
	OtpHitCount         int       `gorm:"default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
