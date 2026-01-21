package models

import "time"

type WalletTransaction struct {
	ID uint `gorm:"primaryKey"`

	UserID    uint
	Reference string `gorm:"index"` // phone stored here
	Amount    float64
	Type      string

	CreatedAt time.Time
}