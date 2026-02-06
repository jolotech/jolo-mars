package models


import (
	"time"

	"gorm.io/gorm"
	"gorm.io/datatypes"

)

// type Admin struct {
// 	ID        uint           `gorm:"primaryKey" json:"-"`
// 	PublicID  string         `gorm:"type:char(15);uniqueIndex;not null" json:"public_id"`
// 	Name      string         `json:"name" gorm:"not null"`
// 	Email     string         `json:"email" gorm:"unique;not null"`
// 	Password  string         `json:"-" gorm:"not null"`
// 	Role      string         `json:"role" gorm:"not null;default:'super-admin'"`
// 	Actions   datatypes.JSON `json:"actions,omitempty" gorm:"type:json"`
// 	CreatedAt time.Time      `json:"createdAt"`
// 	UpdatedAt time.Time      `json:"updatedAt"`
// 	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
// }


// import (
// 	"time"
// 	"gorm.io/gorm"
// 	"gorm.io/datatypes"
// )

type Admin struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	PublicID  string         `gorm:"type:char(15);uniqueIndex;not null" json:"public_id"`

	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`

	Role    string         `json:"role" gorm:"not null;default:'super-admin'"`
	Actions datatypes.JSON `json:"actions,omitempty" gorm:"type:json"`

	// Auth hardening
	MustChangePassword bool       `json:"-" gorm:"not null;default:true"`
	PasswordChangedAt  *time.Time `json:"-"`

	TwoFAEnabled     bool       `json:"-" gorm:"not null;default:false"`
	TwoFASecretEnc   string     `json:"-" gorm:"type:text"` // encrypted secret
	TwoFAConfirmedAt *time.Time `json:"-"`

	FailedLoginAttempts int        `json:"-" gorm:"not null;default:0"`
	LockedUntil         *time.Time `json:"-"`

	LastLoginAt *time.Time `json:"-"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if a.PublicID == "" {
		a.PublicID = GeneratePublicID()
	}
	return nil
}


// func (g *Admin) BeforeCreate(tx *gorm.DB) (err error) {
// 	if g.PublicID == "" {
// 		g.PublicID = GeneratePublicID() 
// 	}
// 	return nil
// }