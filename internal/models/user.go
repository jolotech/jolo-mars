package models


import "time"


type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	FName            string    `json:"f_name" gorm:"column:f_name"`
	LName            string    `json:"l_name" gorm:"column:l_name"`
	Email            string    `json:"email" gorm:"unique"`
	Phone            string    `json:"phone" gorm:"unique"`
	// Password         string    `json:"-" gorm:"not null"`
	Password         string   `json:"-" gorm:"column:password"`
	PasswordHash     string   `json:"-" gorm:"column:password_hash"`
	RefBy            *uint     `json:"ref_by" gorm:"column:ref_by"`
	// RefCode          string    `json:"ref_code" gorm:"column:ref_code"`
	RefCode          string `json:"ref_code" gorm:"type:varchar(100);uniqueIndex"`
	Status           bool      `json:"status" gorm:"default:true"`
	IsPhoneVerified  bool      `json:"is_phone_verified" gorm:"column:is_phone_verified;default:false"`
	IsEmailVerified  bool      `json:"is_email_verified" gorm:"column:is_email_verified;default:false"`
	CMFirebaseToken  *string   `json:"cm_firebase_token" gorm:"column:cm_firebase_token"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
