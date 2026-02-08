package admin_repository

import (
	"errors"

	"github.com/jolotech/jolo-mars/internal/models"
	"gorm.io/gorm"
)

type AdminAthRepo struct {
	db *gorm.DB
}

func NewAdminAuthRepo(db *gorm.DB) *AdminAthRepo {
	return &AdminAthRepo{db: db}
}

func (r *AdminAthRepo) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Admin{}).
		Where("email = ?", email).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}


func (r *AdminAthRepo) CreateAdmin(a *models.Admin) error {
	return r.db.Create(a).Error
}

func (r *AdminAthRepo) GetByEmail(email string) (*models.Admin, error) {
	var a models.Admin
	err := r.db.Where("email = ?", email).First(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &a, err
}
