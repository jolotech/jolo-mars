package user_repository

import (
	// "encoding/json"
	"time"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
)

type Main struct {
	db *gorm.DB
}	


func NewUserMainRepository(db *gorm.DB) *Main {
	return &Main{db: db}
}


func (r *Main) UpdateUser(user *models.User) error {
	// Save updated user to database
    user.UpdatedAt = time.Now()
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
