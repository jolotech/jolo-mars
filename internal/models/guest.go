package models

import (
	"time"

	// "github.com/jolotech/jolo-mars/internal/utils"
	"gorm.io/gorm"
	"crypto/rand"
	"encoding/base64"

	// "github.com/google/uuid"
)

type Guest struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	PublicID  string    `gorm:"type:char(15);uniqueIndex;not null" json:"guest_id"`
	IPAddress string    `gorm:"type:varchar(64);index" json:"ip_address"`
	FCMToken  string    `gorm:"type:text" json:"fcm_token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (g *Guest) BeforeCreate(tx *gorm.DB) (err error) {
	if g.PublicID == "" {
		g.PublicID = utils.GeneratePublicID() // returns 15 chars
	}
	return nil
}


import (
	"crypto/rand"
	"encoding/base64"
)

// GeneratePublicID returns a URL-safe random string of exactly 15 characters
func GeneratePublicID() string {
	b := make([]byte, 11) // 11 bytes ≈ 15 chars base64
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // extremely rare; acceptable for ID generation
	}

	id := base64.RawURLEncoding.EncodeToString(b)
	return id[:15]
}