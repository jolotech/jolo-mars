package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/user"
	"github.com/jolotech/jolo-mars/internal/utils"
	"github.com/jolotech/jolo-mars/types"
)

type UserAuthService struct {
	authRepo *repository.UserAuthRepository
	mainRepo *repository.UserMainRepository
	DB   *gorm.DB
}

// type RegisterRequest struct {
// 	Name     string `json:"name" binding:"required"`
// 	Email    string `json:"email"`
// 	Phone    string `json:"phone" binding:"required"`
// 	Password string `json:"password" binding:"required,min=8"`
// 	RefCode  string `json:"ref_code"`
// }


func NewAuthService(authRepo *repository.UserAuthRepository, mainRepo *repository.UserMainRepository, db *gorm.DB) *UserAuthService {
	return &UserAuthService{
		authRepo: authRepo,
		mainRepo: mainRepo,
		DB: db,
	}
}

func (s *UserAuthService) Register(c *gin.Context, req types.RegisterRequest) (string, interface{}, int, error) {

	// ================= VALIDATION =================
	if msg := utils.ValidateUserRegister(req, s.DB); msg != "" {
		return msg, nil, http.StatusForbidden, errors.New(msg)
	}

	// ================= NAME SPLIT =================
	nameParts := strings.SplitN(req.Name, " ", 2)
	firstName := nameParts[0]
	lastName := ""
	if len(nameParts) > 1 {
		lastName = nameParts[1]
	}

	// ================= REFERRAL LOGIC =================
	var refBy *uint

	if req.RefCode != "" {
		refStatus := repository.GetBusinessSetting(s.DB, "ref_earning_status")
		if refStatus != "1" {
			return utils.ErrorResponse("ref_code", utils.Translate("messages.referer_disable")), 403
		}

		referer, err := repository.FindUserByRefCode(s.DB, req.RefCode)
		if err != nil || !referer.Status {
			return utils.ErrorResponse("ref_code", utils.Translate("messages.referer_code_not_found")), 405
		}

		if repository.IsWalletReferenceUsed(s.DB, req.Phone) {
			return utils.ErrorResponse("phone", utils.Translate("Referrer code already used")), 203
		}

		notification := map[string]interface{}{
			"title":       utils.Translate("messages.Your_referral_code_is_used_by") + " " + firstName + " " + lastName,
			"description": utils.Translate("Be prepare to receive when they complete there first purchase"),
			"order_id":    1,
			"image":       "",
			"type":        "referral_code",
		}

		if utils.GetNotificationStatusData("customer", "customer_new_referral_join", "push_notification_status") &&
			referer.CMFirebaseToken != nil {

			utils.SendPushNotifToDevice(*referer.CMFirebaseToken, notification)
			repository.CreateUserNotification(s.DB, referer.ID, notification)
		}

		refBy = &referer.ID
	}

	// ================= CREATE USER =================
	user := models.User{
		FName:    firstName,
		LName:    lastName,
		Email:    req.Email,
		Phone:    req.Phone,
		RefBy:    refBy,
		Password: utils.HashPassword(req.Password),
		Status:   true,
	}

	newUSer, err := s.authRepo.CreateUser(&user); 
	if err != nil {
		return "error creating user", nil, http.StatusInternalServerError, err
	}

	newUSer.RefCode = utils.GenerateRefererCode(user)
	if err := s.mainRepo.UpdateUser(newUSer); err != nil {
		return "error generating referer code", nil, http.StatusInternalServerError, err
	}

	// ================= TOKEN =================
	token := utils.GenerateAuthToken(user.ID)

	// ================= SETTINGS =================
	loginSettings := repository.GetLoginSettings(s.DB)
	firebaseOTP := repository.GetBusinessSetting(s.DB, "firebase_otp_verification")

	isPhoneVerified := 1
	isEmailVerified := 1

	// ================= PHONE OTP =================
	if loginSettings["phone_verification_status"] == "1" {
		isPhoneVerified = 0

		if firebaseOTP != "1" {
			lastOTP := repository.GetPhoneVerification(s.DB, req.Phone)
			if lastOTP != nil && time.Since(lastOTP.UpdatedAt).Seconds() < 60 {
				wait := 60 - int(time.Since(lastOTP.UpdatedAt).Seconds())
				return utils.OTPWaitError(wait), 405
			}

			otp := utils.GenerateOTP()
			repository.UpsertPhoneOTP(s.DB, req.Phone, otp)

			if !utils.SendSMS(req.Phone, otp) {
				return utils.ErrorResponse("otp", utils.Translate("messages.failed_to_send_sms")), 405
			}

			token = ""
		}
	}

	// ================= EMAIL OTP =================
	if loginSettings["email_verification_status"] == "1" {
		isEmailVerified = 0
		otp := utils.GenerateOTP()
		repository.UpsertEmailOTP(s.DB, req.Email, otp)

		if !utils.SendEmailOTP(req.Email, otp, req.Name) {
			return utils.ErrorResponse("otp", utils.Translate("messages.failed_to_send_mail")), 405
		}

		token = ""
	}

	// ================= REGISTRATION MAIL =================
	utils.SendRegistrationMailIfEnabled(req.Email, req.Name)

	data :=  map[string]interface{}{
		"token":               token,
		"is_phone_verified":   isPhoneVerified,
		"is_email_verified":   isEmailVerified,
		"is_personal_info":    1,
		"is_exist_user":       nil,
		"login_type":          "manual",
		"email":               user.Email,
	}

	return "", data, http.StatusOK, nil
}
