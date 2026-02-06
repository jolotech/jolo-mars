package user_repository

import (
	// "encoding/json"
	"time"
	"errors"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	guest_repo "github.com/jolotech/jolo-mars/internal/repository/guest"

)

type Main struct {
	guestRepo *guest_repo.GuestRepo
	db *gorm.DB
}	


func NewUserMainRepository(db *gorm.DB, GuestRepo *guest_repo.GuestRepo) *Main {
	return &Main{db: db, guestRepo: GuestRepo}
}


func (r *Main) UpdateUser(user *models.User) error {
	// Save updated user to database
    user.UpdatedAt = time.Now()
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Main) GetByEmailOrPhone(email, phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR phone = ?", email, phone).First(&user).Error
	if err != nil {
	    if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil
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



// func MergeGuestCart(db *gorm.DB, userID uint, guestID string) bool {
// 	if guestID == "" || userID == 0 {
// 		return true
// 	}

// 	var guestCartExists bool
// 	db.Model(&models.Cart{}).
// 		Where("user_id = ? AND is_guest = ?", guestID, true).
// 		Select("count(1) > 0").
// 		Scan(&guestCartExists)

// 	if guestCartExists {
// 		db.Where("user_id = ? AND is_guest = ?", userID, false).
// 			Delete(&models.Cart{})
// 	}

// 	db.Model(&models.Cart{}).
// 		Where("user_id = ?", guestID).
// 		Updates(map[string]interface{}{
// 			"user_id":  userID,
// 			"is_guest": false,
// 		})

// 	return true
// }

func (r *Main) MergeGuestCart(db *gorm.DB, userID uint, guestID string) error {
	if guestID == "" || userID == 0 {
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var guestCartExists bool

		if err := tx.Model(&models.Cart{}).
			Where("guest_id = ? AND is_guest = true", guestID).
			Select("count(1) > 0").
			Scan(&guestCartExists).Error; err != nil {
			return err
		}

		if guestCartExists {
			if err := tx.Where("user_id = ?", userID).
				Delete(&models.Cart{}).Error; err != nil {
				return err
			}

			if err := tx.Model(&models.Cart{}).
				Where("guest_id = ? AND is_guest = true", guestID).
				Updates(map[string]interface{}{
					"user_id":  userID,
					"guest_id": gorm.Expr("NULL"),
					"is_guest": false,
				}).Error; err != nil {
				return err
			}
		}

		// Only delete the guest record if merge succeeded or if no cart exists (your choice)
		if err := r.guestRepo.DeleteGuest(tx, guestID); err != nil {
			return err
		}

		return nil
	})
}
