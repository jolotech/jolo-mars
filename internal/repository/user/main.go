package repository

import (
	// "encoding/json"
	"time"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
)

type UserMainRepository struct {
	db *gorm.DB
}	


func NewUserMainRepository(db *gorm.DB) *UserMainRepository {
	return &UserMainRepository{db: db}
}


func (r *UserMainRepository) UpdateUser(user *models.User) error {
	// Save updated user to database
    user.UpdatedAt = time.Now()
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
