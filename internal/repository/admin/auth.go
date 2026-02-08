package admin_repository

import (
	"errors"

	"github.com/jolotech/jolo-mars/internal/models"
	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

func (r *AdminRepo) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Admin{}).
		Where("email = ?", email).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}


func (r *AdminRepo) CreateAdmin(a *models.Admin) error {
	return r.db.Create(a).Error
}

func (r *AdminRepo) GetByEmail(email string) (*models.Admin, error) {
	var a models.Admin
	err := r.db.Where("email = ?", email).First(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &a, err
}
