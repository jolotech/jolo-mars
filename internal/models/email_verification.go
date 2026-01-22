package models

import "time"

type EmailVerification struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;size:255;not null"`
	Token     string    `gorm:"size:6;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
