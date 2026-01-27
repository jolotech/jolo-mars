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

func (r *Main) GetByEmailOrPhone(identifier string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR phone = ?", identifier, identifier).First(&user).Error
	return &user, err
}

func (r *Main) IsWalletReferenceUsed(reference string) bool {
	var tx models.WalletTransaction
	err := r.db.Where("reference = ?", reference).First(&tx).Error
	return err == nil
}


func (r *Main) GetByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	return &user, err
}



func MergeGuestCart(db *gorm.DB, userID uint, guestID string) bool {
	if guestID == "" || userID == 0 {
		return true
	}

	var guestCartExists bool
	db.Model(&models.Cart{}).
		Where("user_id = ? AND is_guest = ?", guestID, true).
		Select("count(1) > 0").
		Scan(&guestCartExists)

	if guestCartExists {
		db.Where("user_id = ? AND is_guest = ?", userID, false).
			Delete(&models.Cart{})
	}

	db.Model(&models.Cart{}).
		Where("user_id = ?", guestID).
		Updates(map[string]interface{}{
			"user_id":  userID,
			"is_guest": false,
		})

	return true
}


 
