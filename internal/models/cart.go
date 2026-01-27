package models

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	IsGuest   bool      `gorm:"default:false"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}