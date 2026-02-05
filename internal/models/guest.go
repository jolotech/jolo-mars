package models

import (
	"time"

	"github.com/google/uuid"
)

type Guest struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	PublicID  uuid.UUID `gorm:"type:char(15);uniqueIndex;not null" json:"public_id"`
	IPAddress string    `gorm:"type:varchar(64);index" json:"ip_address"`
	FCMToken  string    `gorm:"type:text" json:"fcm_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
