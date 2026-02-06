package models

import (
	"time"

	// "github.com/jolotech/jolo-mars/internal/utils"
	"gorm.io/gorm"

	// "github.com/google/uuid"
)

type Guest struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	PublicID  string    `gorm:"type:char(15);uniqueIndex;not null" json:"public_id "`
	IPAddress string    `gorm:"type:varchar(64);index" json:"ip_address"`
	FCMToken  string    `gorm:"type:text" json:"fcm_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (g *Guest) BeforeCreate(tx *gorm.DB) (err error) {
	if g.PublicID == "" {
		g.PublicID = GeneratePublicID() // returns 15 chars
	}
	return nil
}


