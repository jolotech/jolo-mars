package models


import "time"

type User struct {
	ID               uint      `gorm:"primaryKey"`
	FName            string    `gorm:"column:f_name"`
	LName            string    `gorm:"column:l_name"`
	Email            string    `gorm:"unique"`
	Phone            string    `gorm:"unique"`
	Password         string
	RefBy            *uint     `gorm:"column:ref_by"`
	RefCode          string    `gorm:"column:ref_code"`
	Status           bool
	CMFirebaseToken  *string   `gorm:"column:cm_firebase_token"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
