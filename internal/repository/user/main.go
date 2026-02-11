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

func (r *Main) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *Main) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
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
		// ===== LOAD GUESTT CARTS ITEMS ======
		var guestItems []models.Cart
		if err := tx.
			Where("guest_id = ? AND is_guest = true", guestID).
			Find(&guestItems).Error; err != nil {
			return err
		}

		// ====== NOTHING TO MEARGE DELETE GUEST =========
		if len(guestItems) == 0 {
			return r.guestRepo.DeleteGuest(tx, guestID)
			// return nil
		}

		for _, gItem := range guestItems {
			// ======= CHECK IF USER AREADY HAS THIS PRODUCT ===========
			var userItem models.Cart
			err := tx.
				Where("user_id = ? AND is_guest = false AND product_id = ?", userID, gItem.ProductID).
				First(&userItem).Error

			if err == nil {
				// ====== USER HAS IT -> ADD QUANTITY ==============
				if err := tx.Model(&models.Cart{}).
					Where("id = ?", userItem.ID).
					UpdateColumn("quantity", gorm.Expr("quantity + ?", gItem.Quantity)).Error; err != nil {
					return err
				}

				// ============= DELETE GUEST ROW AFTER MARGING QUANTITY
				if err := tx.Delete(&models.Cart{}, gItem.ID).Error; err != nil {
					return err
				}
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				// ===== UUSER DOESN'T HAVE IT -> MOVE GUEST ROW TO USER =======
				if err := tx.Model(&models.Cart{}).
					Where("id = ?", gItem.ID).
					Updates(map[string]interface{}{
						"user_id":  userID,
						"guest_id": gorm.Expr("NULL"),
						"is_guest": false,
					}).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		
		if err := r.guestRepo.DeleteGuest(tx, guestID); err != nil {
			return err
		}

		return nil
	})
}
