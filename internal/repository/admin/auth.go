package admin_repository

import (
	"errors"
	"time"

	"github.com/jolotech/jolo-mars/internal/models"
	"gorm.io/gorm"
)

type AdminAuthRepo struct {
	db *gorm.DB
}

func NewAdminAuthRepo(db *gorm.DB) *AdminAuthRepo {
	return &AdminAuthRepo{db: db}
}


func (r *AdminAuthRepo) UpdatePassword(adminID uint, newHash string, mustChange bool) error {
	return r.db.Model(&models.Admin{}).
		Where("id = ?", adminID).
		Updates(map[string]any{
			"password":             newHash,
			"must_change_password": mustChange,
		}).Error
}


func (r *AdminAuthRepo) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Admin{}).
		Where("email = ?", email).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}


func (r *AdminAuthRepo) CreateAdmin(a *models.Admin) error {
	return r.db.Create(a).Error
}

func (r *AdminAuthRepo) GetByEmail(email string) (*models.Admin, error) {
	var a models.Admin
	err := r.db.Where("email = ?", email).First(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &a, err
}

func (r *AdminAuthRepo) GetByID(id uint) (*models.Admin, error) {
	var a models.Admin
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AdminAuthRepo) Save2FASecret(id uint, encSecret string) error {
	return r.db.Model(&models.Admin{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"two_fa_secret_enc":   encSecret,
			"two_fa_enabled":      false,
			"two_fa_confirmed_at": nil,
		}).Error
}

func (r *AdminAuthRepo) Enable2FA(id uint) error {
	now := time.Now()
	return r.db.Model(&models.Admin{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"two_fa_enabled":      true,
			"two_fa_confirmed_at": &now,
		}).Error
}