package models

import "time"

type UserNotification struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"index"`
	Data   string `gorm:"type:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
