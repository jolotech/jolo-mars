package models

import (
	"time"
	// "github.com/google/uuid"

)

type Cart struct {
	// ID        uint      `gorm:"primaryKey"`
	ID        uint      `gorm:"primaryKey" json:"-"`
	// PublicID  uuid.UUID `gorm:"type:char(36);uniqueIndex;not null" json:"public_id"`
	UserID    uint      `gorm:"not null"`
	IsGuest   bool      `gorm:"default:false"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}