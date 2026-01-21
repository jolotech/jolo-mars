package models


import "time"

type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	FName            string    `json:"f_name" gorm:"column:f_name"`
	LName            string    `json:"l_name" gorm:"column:l_name"`
	Email            string    `json:"email" gorm:"unique"`
	Phone            string    `json:"phone" gorm:"unique"`
	Password         string    `json:"-" gorm:"not null"`
	RefBy            *uint     `json:"ref_by" gorm:"column:ref_by"`
	RefCode          string    `json:"ref_code" gorm:"column:ref_code"`
	Status           bool      `json:"status" gorm:"default:true"`
	CMFirebaseToken  *string   `json:"cm_firebase_token" gorm:"column:cm_firebase_token"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
