package models

import "time"

type Guest struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IPAddress string    `gorm:"type:varchar(64);index" json:"ip_address"`
	FCMToken  string    `gorm:"type:text" json:"fcm_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
