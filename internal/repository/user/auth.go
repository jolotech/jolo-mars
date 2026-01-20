package repository


import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"

)


type UserAuthRepository struct {
	db *gorm.DB
}	

func NewUserAuthRepository(db *gorm.DB) *UserAuthRepository {
	return &UserAuthRepository{db: db}
}

func FindUserByRefCode(db *gorm.DB, code string) (*models.User, error) {
	var user models.User
	err := db.Where("ref_code = ?", code).First(&user).Error
	return &user, err
}

func CreateUserNotification(db *gorm.DB, userID uint, data map[string]interface{}) {
	payload, _ := json.Marshal(data)
	db.Table("user_notifications").Create(map[string]interface{}{
		"user_id":    userID,
		"data":       payload,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})
}
