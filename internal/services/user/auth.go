package services

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/user"
	"github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/utils"
	"github.com/jolotech/jolo-mars/internal/helpers/notifications"
	"github.com/jolotech/jolo-mars/internal/helpers/verifications"
	"github.com/jolotech/jolo-mars/types"
)

type UserAuthService struct {
	authRepo *user_repository.Auth
	usermainRepo *user_repository.Main
	adminmainRepo *admin_repository.Main
	DB   *gorm.DB
}

// type RegisterRequest struct {
// 	Name     string `json:"name" binding:"required"`
// 	Email    string `json:"email"`
// 	Phone    string `json:"phone" binding:"required"`
// 	Password string `json:"password" binding:"required,min=8"`
// 	RefCode  string `json:"ref_code"`
// }


func NewAuthService(authRepo *user_repository.Auth, usermainRepo *user_repository.Main, adminmainRepo *admin_repository.Main, db *gorm.DB) *UserAuthService {
	return &UserAuthService{
		authRepo: authRepo,
		usermainRepo: usermainRepo,
		adminmainRepo: adminmainRepo,
		DB: db,
	}
}

func (s *UserAuthService) Register(c *gin.Context, req types.RegisterRequest) (string, interface{}, int, error) {

	// ================= VALIDATION =================
	if msg := utils.ValidateUserRegister(req, s.DB); msg != "" {
		return msg, nil, http.StatusForbidden, errors.New("Validation error")
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
		refStatus := s.adminmainRepo.GetBusinessSetting("ref_earning_status").(bool)
		// refStatus := setting.(bool)
		if  !refStatus {
			return "referer not available", nil, http.StatusForbidden, errors.New("referer not available")
		}

		referer, err := s.authRepo.FindUserByRefCode(req.RefCode)
		if err != nil || !referer.Status {
			return "invalid referer code", nil, http.StatusNotFound, errors.New("invalid referer code")
		}

		if s.usermainRepo.IsWalletReferenceUsed(req.Phone) {
			return "Referrer code already used", nil, http.StatusForbidden, errors.New("Referrer code already used")
		}

		notification := map[string]interface{}{
			"title":       "Your referral code is used by" + " " + firstName + " " + lastName,
			"description": "Be prepare to receive a coupon when they complete there first purchase",
			"order_id":    1,
			"image":       "",
			"type":        "referral_code",
		}

		if helpers.GetNotificationStatusData("customer", "customer_new_referral_join", "push_notification_status", nil) &&
			referer.CMFirebaseToken != nil {

			helpers.SendPushNotifToDevice(*referer.CMFirebaseToken, notification)
			s.authRepo.CreateUserNotification(referer.ID, notification)
		}

		refBy = &referer.ID
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
	    return "", nil, http.StatusInternalServerError, err
	}


	// ================= CREATE USER =================
	user := models.User{
		FName:    firstName,
		LName:    lastName,
		Email:    req.Email,
		Phone:    req.Phone,
		PasswordHash: hashedPassword,
		RefBy:    refBy,
		Password: hashedPassword,
		Status:   true,
	}

	newUSer, err := s.authRepo.CreateUser(&user); 
	if err != nil {
		return "error creating user", nil, http.StatusInternalServerError, err
	}

	newUSer.RefCode = utils.GenerateRefererCode(s.DB)
	if err := s.usermainRepo.UpdateUser(newUSer); err != nil {
		return "error generating referer code", nil, http.StatusInternalServerError, err
	}

	// ================= TOKEN =================
	token, _ := utils.GenerateAuthToken(user.Email, user.ID)

	// ================= SETTINGS =================
	loginSettings := s.adminmainRepo.GetLoginSettings()
	firebaseOTP := s.adminmainRepo.GetBusinessSetting("firebase_otp_verification").(bool)

	isPhoneVerified := false
	isEmailVerified := false


	// ================= PHONE OTP =================
	if loginSettings.PhoneVerification {

	    if !firebaseOTP {

		    lastOTP, _ := user_repository.GetVerification(s.DB, req.Phone)
		    if lastOTP != nil {
			    if user_repository.IsOtpLocked(lastOTP) {
				    return "too many attempts", nil, 403, errors.New("otp locked")
			    }
			    ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
			    if !ok {
				    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
			    }
		    }

		    otp := utils.GenerateOTP()
		    user_repository.UpsertOTP(s.DB, req.Phone, otp)

		    if !otp_helpers.SendSMS(req.Phone, otp) {
			    return "failed to send sms", nil, 405, errors.New("sms failed")
		    }
		    user_repository.IncrementOtpHit(s.DB, req.Phone)

		    token = ""
	    }
    }


    if loginSettings.EmailVerification {
	    isEmailVerified = false

	    lastOTP, _ := user_repository.GetVerification(s.DB, req.Email)
	    if lastOTP != nil {
		    if user_repository.IsOtpLocked(lastOTP) {
				return "too many attempts", nil, 403, errors.New("otp locked")
			}
		    ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
		    if !ok {
			    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
		    }
		}

		otp := utils.GenerateOTP()
	    user_repository.UpsertOTP(s.DB, req.Email, otp)

	    if !otp_helpers.SendEmailOTP(req.Email, otp, req.Name) {
		    return "failed_to_send_mail", nil, 405, errors.New("mail failed")
	    }
	    user_repository.IncrementOtpHit(s.DB, req.Email)

	    token = ""
	}


		// ================= REGISTRATION MAIL =================
	// utils.SendRegistrationMailIfEnabled(req.Email, req.Name)

	data :=  map[string]interface{}{
		"token":               token,
		"is_phone_verified":   isPhoneVerified,
		"is_email_verified":   isEmailVerified,
		"is_personal_info":    1,
		"is_exist_user":       nil,
		"email":               user.Email,
	}

	return "", data, http.StatusOK, nil
}
