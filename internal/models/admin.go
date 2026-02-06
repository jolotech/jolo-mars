package models


import (
	"time"

	"gorm.io/gorm"
	"gorm.io/datatypes"
	// "github.com/google/uuid"

)

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	PublicID  string         `gorm:"type:char(15);uniqueIndex;not null" json:"public_id"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"-" gorm:"not null"`
	Role      string         `json:"role" gorm:"not null;default:'super-admin'"`
	Actions   datatypes.JSON `json:"actions,omitempty" gorm:"type:json"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}


func (g *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if g.PublicID == "" {
		g.PublicID = GeneratePublicID() 
	}
	return nil
}