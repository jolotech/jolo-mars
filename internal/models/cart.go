package models

import (
	"time"
	// "github.com/google/uuid"
	"gorm.io/gorm"

)

// type Cart struct {
// 	ID        uint      `gorm:"primaryKey" json:"-"`
// 	PublicID  string    `gorm:"type:char(15);uniqueIndex;not null" json:"public_id"`
// 	UserID    uint      `gorm:"not null"`
// 	IsGuest   bool      `gorm:"default:false"`
// 	GuestId   string    `gorm:"nulable"`
// 	ProductID uint      `gorm:"not null"`
// 	Quantity  int       `gorm:"default:1"`
// 	CreatedAt time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime"`
// }

type Cart struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	PublicID string `gorm:"type:char(15);uniqueIndex;not null" json:"cart_id"`

	UserID  *uint   `gorm:"index" json:"-"`                             
	GuestID *string `gorm:"type:char(15);index" json:"-"`               
	IsGuest bool    `gorm:"default:false;index" json:"is_guest"`

	ProductID uint `gorm:"not null;index" json:"product_id"`
	Quantity  int  `gorm:"not null;default:1" json:"quantity"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}


func (g *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	if g.PublicID == "" {
		g.PublicID = GeneratePublicID() // returns 15 chars
	}
	return nil
}