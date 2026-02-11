package services

import (
	"errors"
	"net/http"
	"strings"

	// "log"
	// "time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	// "github.com/jolotech/jolo-mars/config"
	"github.com/jolotech/jolo-mars/internal/helpers/email"
	"github.com/jolotech/jolo-mars/internal/helpers/notifications"
	"github.com/jolotech/jolo-mars/internal/helpers/verifications"
	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/repository/user"
	"github.com/jolotech/jolo-mars/internal/utils"
	"github.com/jolotech/jolo-mars/types"
	// "github.com/google/uuid"
)


type UserAuthService struct {
	authRepo *user_repository.Auth
	usermainRepo *user_repository.Main
	adminmainRepo *admin_repository.Main
	DB   *gorm.DB
}

type OTPSendFunc func(otp string) error

func NewAuthService(authRepo *user_repository.Auth, usermainRepo *user_repository.Main, adminmainRepo *admin_repository.Main, db *gorm.DB) *UserAuthService {
	return &UserAuthService{authRepo, usermainRepo, adminmainRepo, db,}
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
	}

	// ================= OTP SETTINGS =================
	loginSettings := s.adminmainRepo.GetLoginSettings()
	firebaseOTP := s.adminmainRepo.GetBusinessSetting("firebase_otp_verification").(bool)
	phoneOption := loginSettings.PhoneVerification
	emailOption := loginSettings.EmailVerification


	// ================= OTP OPTION CHECK =================
	if req.OtpOption == "phone" && !phoneOption {
		return "phone otp not enabled", nil, http.StatusForbidden, errors.New("phone otp not enabled")
	}
	if req.OtpOption == "email" && !emailOption {
		return "email otp not enabled", nil, http.StatusForbidden, errors.New("email otp not enabled")
	}
	

	// ================= PHONE OTP =================
	if phoneOption && req.OtpOption == "phone" {

	    if !firebaseOTP {
			// ================= SAVE USER AND SEND OTP SMS =================
			s.authRepo.SaveSignUpUSer(user)
			return HandleOTP(s.DB, req.Phone, func(otp string) error {
				if !otp_helpers.SendSMS(req.Phone, otp) {
				    return errors.New("sms failed")
					// return "failed to send sms",  405, errors.New("sms failed")
			    }
			    return nil
		    },)
	    }
    }

	// ================= SAVE USER AND SEND OTP EMAIL =================
    if emailOption && req.OtpOption == "email" {
		s.authRepo.SaveSignUpUSer(user)
		return HandleOTP(s.DB, req.Email, func(otp string) error {
			return email.SendEmail(otp, &user).Verification()
		},)
	}


	// ================= RESPONSE =================
	if emailOption && req.OtpOption == "email" {
		return "verification email sent", nil, http.StatusOK, nil
	}
	if phoneOption && req.OtpOption == "phone" {
		return "verification sms sent", nil, http.StatusOK, nil
	}

	return "registration successful", nil, http.StatusOK, nil
}

// ================= VERIFY OTP =================

func (s *UserAuthService) VerifyOTP(req types.VerifyOTPRequest) (string, any, int, error) {

	var isPhone = req.VerificationMethod == "phone"
    var identifier string

	// =====================GET USER =====================
	user, err := s.usermainRepo.GetByEmailOrPhone(req.Email, req.Phone)
	if err != nil || user == nil {
		return "User not found", nil, http.StatusNotFound, err
	}

	//================== OTP CHECK =======================
	if isPhone {
		identifier = req.Phone
	}else {
		identifier = req.Email
	}
	msg, verification, statusCode, err := user_repository.OTPCheck(s.DB, identifier, req.OTP)
	if err != nil || verification == nil {
		return msg, verification, statusCode, err
	}

	// ===================== UPDATE USER STATUS AND DELETE VERIFICATION =================
	if isPhone {
		user.IsPhoneVerified = true
		user_repository.DeleteVerification(s.DB, identifier, req.OTP)
	} else {
		user.IsEmailVerified = true
		user_repository.DeleteVerification(s.DB, identifier, req.OTP)
	}

	// ================= TOKEN =================
	token, _ := utils.GenerateAuthToken(user.Email, user.ID)

	if !user.Status {
		email.SendEmail(nil, user).Welcome()
	}

	user.Status = true
	if err := s.usermainRepo.UpdateUser(user); err != nil {
		return "Failed to verify user", nil, http.StatusInternalServerError, err
	}

	// ================= MERGE GUEST CART =================
	if req.GuestID != nil {
		s.usermainRepo.MergeGuestCart(s.DB, user.ID, *req.GuestID)
	}

	// ================= RESPONSE =================
	// data := map[string]interface{}{
	// 	"user":              user,
	// 	"token":             token,
	// }

	data := &types.AuthLoginData{
		User:              user,
		Token:             token,
	}

	return "Verification successful", data, http.StatusOK, nil
}


func (s *UserAuthService) ResendOTP(req types.ResendOTPRequest) (string, any, int, error){

	var user *models.User
	var err error

	isPhone := req.VerificationMethod == "phone"
	isEmail := req.VerificationMethod == "email"

	// =====================GET USER =====================
	user, err = s.usermainRepo.GetByEmailOrPhone(req.Email, req.Phone)
	if err != nil || user == nil {
		return "User not found", nil, http.StatusNotFound, err
	}
		
	// ================= OTP SETTINGS =================
	loginSettings := s.adminmainRepo.GetLoginSettings()
	firebaseOTP := s.adminmainRepo.GetBusinessSetting("firebase_otp_verification").(bool)
	phoneOption := loginSettings.PhoneVerification
	emailOption := loginSettings.EmailVerification


	// ================= OTP OPTION CHECK =================
	if isPhone && !phoneOption {
		return "phone otp not enabled", nil, http.StatusForbidden, errors.New("phone otp not enabled")
	}
	if isEmail && !emailOption {
		return "email otp not enabled", nil, http.StatusForbidden, errors.New("email otp not enabled")
	}

	// ================= PHONE OTP =================
	if phoneOption && isPhone {

	    if !firebaseOTP {
			return HandleOTP(s.DB, req.Phone, func(otp string) error {
				if !otp_helpers.SendSMS(req.Phone, otp) {
				    return errors.New("sms failed")
					// return "failed to send sms",  405, errors.New("sms failed")
			    }
			    return nil
		    },
		)}
    }

	// ================= EMAIL OTP =================
    if emailOption && isEmail {
		return HandleOTP(s.DB, req.Email, func(otp string) error {
			return email.SendEmail(otp, user).Verification()
		},)
	}
	return "OTP Sent Successfully", nil, http.StatusOK, nil
}


func HandleOTP(db *gorm.DB, identifier string, sendOTP OTPSendFunc,) (string, interface{}, int, error) {

	lastOTP, _ := user_repository.GetVerification(db, identifier)
	if lastOTP != nil {

		if !lastOTP.IsActive {
			return "OTP deactivated", nil, http.StatusBadRequest, errors.New("Invalid OTP")
		}

		if user_repository.IsOtpLocked(lastOTP) {
			lastOTP.IsActive = false
			user_repository.UpdateVerification(db, *lastOTP)
			return "too many attempts", nil, http.StatusForbidden, errors.New("otp locked")
		}

		ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
		if !ok {
			return utils.OTPWaitError(wait), nil, http.StatusMethodNotAllowed, errors.New("otp wait error")
		}
	}

	otp := utils.GenerateOTP()
	user_repository.UpsertOTP(db, identifier, otp)

	if err := sendOTP(otp); err != nil {
		return err.Error(), nil, http.StatusInternalServerError, err
	}

	user_repository.IncrementOtpHit(db, identifier)

	return "OTP sent successfully", nil, http.StatusOK, nil
}


func (s *UserAuthService) ForgetPassword(req types.ResendOTPRequest) (string, any, int, error) {

	isPhone := req.VerificationMethod == "phone"
	isEmail := req.VerificationMethod == "email"
	
	user, err := s.usermainRepo.GetByEmailOrPhone(req.Email, req.Phone)
	if err != nil || user == nil{
		return "user not found", nil, http.StatusNotFound, err
	}

	// ================= OTP SETTINGS =================
	loginSettings := s.adminmainRepo.GetLoginSettings()
	firebaseOTP := s.adminmainRepo.GetBusinessSetting("firebase_otp_verification").(bool)
	phoneOption := loginSettings.PhoneVerification
	emailOption := loginSettings.EmailVerification


	// ================= OTP OPTION CHECK =================
	if isPhone && !phoneOption {
		return "phone otp not enabled", nil, http.StatusForbidden, errors.New("phone otp not enabled")
	}
	if isEmail && !emailOption {
		return "email otp not enabled", nil, http.StatusForbidden, errors.New("email otp not enabled")
	}

	
	// ================= OTP OPTION CHECK =================
	if isPhone && !phoneOption {
		return "phone otp not enabled", nil, http.StatusForbidden, errors.New("phone otp not enabled")
	}
	if isEmail && !emailOption {
		return "email otp not enabled", nil, http.StatusForbidden, errors.New("email otp not enabled")
	}

	// ================= PHONE OTP =================
	if phoneOption && isPhone {

	    if !firebaseOTP {
			return HandleOTP(s.DB, req.Phone, func(otp string) error {
				if !otp_helpers.SendSMS(req.Phone, otp) {
				    return errors.New("sms failed")
			    }
			    return nil
		    },
		)}
    }

	// ================= EMAIL OTP =================
    if emailOption && isEmail {
		return HandleOTP(s.DB, req.Email, func(otp string) error {
			return email.SendEmail(otp, user).ForgetPassword()
		},)
	}
	return "OTP Sent Successfully", nil, http.StatusOK, nil
}

func (s *UserAuthService) ResetPassword(req types.ResetPasswordSubmitRequest) (string, any, int, error) {

	var identifier string
	isPhone := req.VerificationMethod == "Phone"

	//==================== GET USER ===============
	user, err := s.usermainRepo.GetByEmailOrPhone(req.Email, req.Phone)
	if err != nil || user == nil {
		return "user not found", nil, http.StatusNotFound, errors.New("not found")
	}

	// ============ Password CHECK ===============
	if req.Password != req.ConfirmPassword {
		return "passwords do not match", nil, http.StatusUnauthorized, errors.New("mismatch")
	}
	isOldPassword := utils.ComparePassword(user.Password, req.Password)
	if isOldPassword {
		return "cant use previous password", nil, http.StatusUnauthorized, errors.New("cant use old password update to a new password")
	}

	//================== OTP CHECK ===================
	if isPhone {
		identifier = req.Phone
	}else {
		identifier = req.Email
	}
	msg, verification, statusCode, err := user_repository.OTPCheck(s.DB, identifier, req.ResetToken)
	if err != nil  {
		return msg, verification, statusCode, err
	}

	//=============== HASH NEW PASSWORD ===================
	hashedPass, err := utils.HashPassword(req.ConfirmPassword)
	if err != nil {
		return "failed to hash password", nil, http.StatusInternalServerError, err
	}

	//=============== SAVE PASSWORD =======================
	user.Password = hashedPass
	if err := s.usermainRepo.UpdateUser(user); err != nil {
		return "failed to update password", nil, http.StatusInternalServerError, err
	}

	//=============== DELETE VERIFICATION AND SEND EMAIL =====================
	// user_repository.DeleteVerification(s.DB, )
	email.SendEmail(nil, user).ResetPassword();

	//================ SUCCESS RESPONSE ====================
	return "Password changed successfully.", nil, http.StatusOK, nil
}


func (s *UserAuthService) Login(req types.UserLoginRequest) (string, any, int, error) {
	// 1) Find user
	// var user *models.User
	var err error

	// switch req.Method {
	// case "email":
	// 	user, err = s.userRepo.FindByEmail(ctx, req.EmailOrPhone)
	// case "phone":
	// 	user, err = s.userRepo.FindByPhone(ctx, req.EmailOrPhone)
	// default:
	// 	return "invalid field_type", nil, http.StatusBadRequest, errors.New("invalid field_type")
	// }

	// =====================GET USER =====================
	user, err := s.usermainRepo.GetByEmailOrPhone(req.Email, req.Phone)
	if err != nil || user == nil {
		return "User credential does not match", nil, http.StatusUnauthorized, err
	}

	isCoorectPassword := utils.ComparePassword(user.Password, req.Password)
	if !isCoorectPassword {
		return "User credential does not match", nil, http.StatusUnauthorized, errors.New("incorect password")
	}

	// 3) Blocked check
	if !user.Status {
		return "your account is temporarily blocked", nil, http.StatusForbidden, errors.New("account blocked")
	}

	// 4) check-ref-code (Ensure ref_code exists)
	// Only generate + update if missing.
	if  user.RefCode == "" {
		refCode := utils.GenerateRefererCode(s.DB)

		// Update only if still empty (safe for concurrency)
		_, upErr := s.usermainRepo.EnsureRefCode(user.ID, refCode)
		if upErr != nil {
			return "failed to update referal code", nil, http.StatusInternalServerError, upErr
		}
	}


	// 7) Token only when personal info exists 
	var token string
	// if isPersonalInfo == 1 {
		tk, tkErr := utils.GenerateAuthToken(user.Email, user.ID)
		if tkErr != nil {
			return "login error", nil, http.StatusInternalServerError, tkErr
		}
		token = tk

		// 8) merge guest cart if guest_id provided
		if req.GuestID != nil && *req.GuestID != "" {
			// if merge fails you can decide:
			// - either fail login
			// - or ignore and still login
			if err := s.usermainRepo.MergeGuestCart(s.DB, user.ID, *req.GuestID); err != nil {
				// I recommend: don't block login, but log error
				// return "failed to merge guest cart", nil, http.StatusInternalServerError, err
			}
		}
	// }

	// 9) response payload (matches your PHP response)
	

	data := &types.AuthLoginData{
		User:              user,
		Token:             token,
	}

	return "login successful", data, http.StatusOK, nil
}