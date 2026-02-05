package guest_repo

import (
	"time"

	"gorm.io/gorm"

	"github.com/jolotech/jolo-mars/internal/models"
)


type GuestRepo struct {
	db *gorm.DB
}

// type Repository interface {
//     CreateGuest(guest *models.Guest) error
// }

func NewGuestRepo(db *gorm.DB) *GuestRepo {
	return &GuestRepo{db: db}
}

func (r *GuestRepo) CreateGuest(guest *models.Guest) error {
	guest.CreatedAt = time.Now()
	guest.UpdatedAt = time.Now()
	if err := r.db.Create(guest).Error; err != nil {
		return err
	}
	return  nil
}


func DeleteVerification(db *gorm.DB, id string) error {
	if err := db.Where("guest_id = ? ", id).
		Delete(&models.Guest{}).Error; err != nil {
		return err
	}
	return nil
}
