package user_repository

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/internal/repository/user"
	"github.com/jolotech/jolo-mars/internal/utils"
	"gorm.io/gorm"
)


type Auth struct {
	db *gorm.DB
	Main *Main
}	

func NewUserAuthRepository(db *gorm.DB, Main *Main) *Auth {
	return &Auth{
		Main: Main,
		db: db,
	}
}


func (r *Auth) CreateUser(user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Auth) CreateGuest(guest *models.Guest) error {
	guest.CreatedAt = time.Now()
	guest.UpdatedAt = time.Now()
	if err := r.db.Create(guest).Error; err != nil {
		return err
	}
	return  nil
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

func(r *Auth) SaveSignUpUSer(user models.User) (string, any, int, error) {
	newUSer, err := r.CreateUser(&user); 
	if err != nil {
		return "error creating user", nil, http.StatusInternalServerError, err
	}
	newUSer.RefCode = utils.GenerateRefererCode(r.db)
	if err := r.Main.UpdateUser(newUSer); err != nil {
		return "error generating referer code", nil, http.StatusInternalServerError, err
	}
	return "", nil, http.StatusOK, nil
}