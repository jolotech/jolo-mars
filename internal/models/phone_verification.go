package models

import "time"

type PhoneVerification struct {
	ID          uint      `gorm:"primaryKey"`
	Phone       string    `gorm:"uniqueIndex;size:20;not null"`
	Token       string    `gorm:"size:10;not null"`
	OtpHitCount int       `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
