package domain

import "time"

type MigrationHistory struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(255);uniqueIndex"`
	ExecutedAt time.Time `gorm:"autoCreateTime"`
}
