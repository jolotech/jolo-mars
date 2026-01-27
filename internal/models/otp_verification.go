package models

import "time"


type OtpVerification struct {
	ID                  uint      `gorm:"primaryKey"`
	Phone               string    `gorm:"uniqueIndex;size:20;not null"`
	Token               string    `gorm:"size:6;not null"`
	VerificationMethod  string    `gorm:"not null"`
	IsActive            bool      `gorm:"default:false"`
	OtpHitCount         int       `gorm:"default:0"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
