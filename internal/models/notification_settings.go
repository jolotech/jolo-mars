package models

import "time"

type NotificationSetting struct {
	ID                        uint `gorm:"primaryKey"`

	Type                      string `gorm:"size:50;index"` // customer, store, admin
	Key                       string `gorm:"size:100;index"`

	PushNotificationStatus    string `gorm:"size:20;default:inactive"`
	EmailNotificationStatus   string `gorm:"size:20;default:inactive"`
	SmsNotificationStatus     string `gorm:"size:20;default:inactive"`

	StoreID                   *uint  `gorm:"index"` // nullable (for store override)

	CreatedAt time.Time
	UpdatedAt time.Time
}
