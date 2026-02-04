package models

import (
	"time"
	"github.com/google/uuid"
)

type WalletTransaction struct {
	// ID uint `gorm:"primaryKey"`
    ID               uint      `gorm:"primaryKey" json:"-"`
	PublicID         uuid.UUID `gorm:"type:char(36);uniqueIndex;not null" json:"public_id"`
	UserID    uint
	Reference string `gorm:"index"` // phone stored here
	Amount    float64
	Type      string
	CreatedAt time.Time
}