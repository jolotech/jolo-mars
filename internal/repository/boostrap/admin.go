package admin_repository

import (

	"github.com/jolotech/jolo-mars/internal/models"
	"gorm.io/gorm"
)

type AdminBoostrap struct {
	db *gorm.DB
}

func NewAdminBoostrapRepository(db *gorm.DB) *AdminBoostrap {
	return &AdminBoostrap{db: db}
}

func (r *AdminBoostrap) AnySuperAdminExists() (bool, error) {
	var count int64
	err := r.db.Model(&models.Admin{}).
		Where("role = ?", "super-admin").
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}

func (r *AdminBoostrap) CreateSuperAdmin(a *models.Admin) error {
	return r.db.Create(a).Error
}
