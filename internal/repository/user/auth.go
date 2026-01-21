package user_repository


import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"

)


type Auth struct {
	db *gorm.DB
}	

func NewUserAuthRepository(db *gorm.DB) *Auth {
	return &Auth{db: db}
}


func (r *Auth) CreateUser(user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Auth) FindUserByRefCode(code string) (*models.User, error) {
	var user models.User
	err := r.db.Where("ref_code = ?", code).First(&user).Error
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
 