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

func (r *Auth) CreateUserNotification(userID uint, data map[string]interface{}) error {

	payload, _ := json.Marshal(data)

	return r.db.Create(&models.UserNotification{
		UserID: userID,
		Data:   string(payload),
	}).Error
}

 