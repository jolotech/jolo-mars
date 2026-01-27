package models
import "time"

type GuestCart struct {
	ID        uint      `gorm:"primaryKey"`
	SessionID string    `gorm:"not null;index"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
