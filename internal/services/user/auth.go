package services

import (
	"errors"
	// "log"
	"net/http"
	"strings"

	// "log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jolotech/jolo-mars/internal/helpers/email"
	"github.com/jolotech/jolo-mars/internal/helpers/notifications"
	"github.com/jolotech/jolo-mars/internal/helpers/verifications"
	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/repository/user"
	"github.com/jolotech/jolo-mars/internal/utils"
	"github.com/jolotech/jolo-mars/types"
)

type UserAuthService struct {
	authRepo *user_repository.Auth
	usermainRepo *user_repository.Main
	adminmainRepo *admin_repository.Main
	DB   *gorm.DB
}

type OTPSendFunc func(otp string) error

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

	// ================= TOKEN =================
	// token, _ := utils.GenerateAuthToken(user.Email, user.ID)



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
			// ================= SAVE USER AND SEND OTP=================
			s.authRepo.SaveSignUpUSer(user)
			return HandleOTP(s.DB, req.Phone, func(otp string) error {
				if !otp_helpers.SendSMS(req.Phone, otp) {
				    return errors.New("sms failed")
					// return "failed to send sms",  405, errors.New("sms failed")
			    }
			    return nil
		    },)

		    // lastOTP, _ := user_repository.GetVerification(s.DB, req.Phone)
		    // if lastOTP != nil {
			//     if user_repository.IsOtpLocked(lastOTP) {
			// 	    return "too many attempts", nil, 403, errors.New("otp locked")
			//     }
			//     ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
			//     if !ok {
			// 	    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
			//     }
		    // }

		    // otp := utils.GenerateOTP()
		    // user_repository.UpsertOTP(s.DB, req.Phone, otp)

		    // if !otp_helpers.SendSMS(req.Phone, otp) {
			//     return "failed to send sms", nil, 405, errors.New("sms failed")
		    // }
		    // user_repository.IncrementOtpHit(s.DB, req.Phone)
	    }
    }

	// ================= EMAIL OTP =================

    if emailOption && req.OtpOption == "email" {
		// ================= SAVE USER AND SEND OTP=================
		s.authRepo.SaveSignUpUSer(user)
		return HandleOTP(s.DB, req.Email, func(otp string) error {
			return email.SendEmail(otp, &user).Verification()
		},)
		
	    // lastOTP, _ := user_repository.GetVerification(s.DB, req.Email)
	    // if lastOTP != nil {
		//     if user_repository.IsOtpLocked(lastOTP) {
		// 		return "too many attempts", nil, 403, errors.New("otp locked")
		// 	}
		//     ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
		//     if !ok {
		// 	    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
		//     }
		// }

		// otp := utils.GenerateOTP()
		// user_repository.UpsertOTP(s.DB, req.Email, otp)

		// err := email.SendEmail(otp, &user).Verification()
		// if err != nil {
		// 	return "failed to send email", nil, http.StatusInternalServerError, err
		// }
		// user_repository.IncrementOtpHit(s.DB, req.Email)
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
	var user *models.User
	var err error

	var isPhone = req.VerificationType == "phone"

	// =====================GET USER =====================
	if isPhone {
		user, err = s.usermainRepo.GetByEmailOrPhone(req.Phone)
	} else {
		user, err = s.usermainRepo.GetByEmailOrPhone(req.Email)
	}

	if err != nil || user == nil {
		return "User not found", nil, http.StatusNotFound, err
	}

	//====================== CHECK OTP MATCH =====================
	// var verification *models.OtpVerification
	// var err  error
	// if isPhone {
	// 	verification, err = user_repository.GetVerification(s.DB, req.Phone)
	// } else {
	// 	verification, err = user_repository.GetVerification(s.DB, req.Email)
	// }

	verification, err := func() (*models.OtpVerification, error) {
		if isPhone {
            return user_repository.GetVerification(s.DB, req.Phone)
		}
        return user_repository.GetVerification(s.DB, req.Email)
	}()

	if err != nil || verification == nil {
		return "Invalid verification", nil, http.StatusUnavailableForLegalReasons, errors.New("Invalid verification")
	}

	if !verification.IsActive {
		return "OTP deactivated", nil, http.StatusBadRequest, errors.New("Invalid OTP")
	}

	// =================== Check FOR EXPIRED OTP ========================

	if user_repository.IsOTPExpired(verification.UpdatedAt, 10*time.Minute) {
		return "OTP expired", nil, http.StatusBadRequest, errors.New("OTP expired")
	}


	// ====================== GET OTP VERIFICATION =====================
	if verification.Token != req.OTP {
		return "OTP does not match", nil, http.StatusBadRequest, errors.New("OTP does not match")
	}

	// ===================== DEACTIVATE OTP ========================
	verification.IsActive = false
	user_repository.UpdateVerification(s.DB, *verification)

	// ===================== UPDATE USER VERIFICATION STATUS =================
	if isPhone {
		user.IsPhoneVerified = true
	} else {
		user.IsEmailVerified = true
	}

	// ================= TOKEN =================
	token, _ := utils.GenerateAuthToken(user.Email, user.ID)

	// ================= MERGE GUEST CART =================
	if req.GuestID != nil {
		user_repository.MergeGuestCart(s.DB, user.ID, *req.GuestID)
	}

	if !user.Status {
		email.SendEmail(nil, user).Welcome()
	}

	user.Status = true

	if err := s.usermainRepo.UpdateUser(user); err != nil {
		return "Failed to verify user", nil, http.StatusInternalServerError, err
	}

	// ================= RESPONSE =================
	data := map[string]interface{}{
		"user":              user,
		"token":             token,
	}

	return "Verification successful", data, http.StatusOK, nil
}


func (s *UserAuthService) ResendOTP(req types.ResendOTPRequest) (string, any, int, error){

	var user *models.User
	var err error

	isPhone := req.VerificationType == "phone"
	isEmail := req.VerificationType == "email"

	// =====================GET USER =====================
	if isPhone {
		user, err = s.usermainRepo.GetByEmailOrPhone(req.Phone)
	} else {
		user, err = s.usermainRepo.GetByEmailOrPhone(req.Email)
	}

	if err != nil || user == nil {
		return "User not found", nil, http.StatusNotFound, err
	}

	//==================== GET VERIFICATION ======================

	// verification, err := func() (*models.OtpVerification, error) {
	// 	if isPhone {
    //         return user_repository.GetVerification(s.DB, req.Phone)
	// 	}
    //     return user_repository.GetVerification(s.DB, req.Email)
	// }()

	// if err != nil || verification == nil {
		
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
		    },)

		//     lastOTP, _ := user_repository.GetVerification(s.DB, req.Phone)
		//     if lastOTP != nil {
		// 	    if user_repository.IsOtpLocked(lastOTP) {
		// 		    return "too many attempts", nil, 403, errors.New("otp locked")
		// 	    }
		// 	    ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
		// 	    if !ok {
		// 		    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
		// 	    }
		//     }

		//     otp := utils.GenerateOTP()
		//     user_repository.UpsertOTP(s.DB, req.Phone, otp)

		//     if !otp_helpers.SendSMS(req.Phone, otp) {
		// 	    return "failed to send sms", nil, 405, errors.New("sms failed")
		//     }
		//     user_repository.IncrementOtpHit(s.DB, req.Phone)
	    }
    }

	// ================= EMAIL OTP =================

    if emailOption && isEmail {

		return HandleOTP(s.DB, req.Email, func(otp string) error {
			return email.SendEmail(otp, user).Verification()
		},)

	    // lastOTP, _ := user_repository.GetVerification(s.DB, req.Email)
	    // if lastOTP != nil {
		//     if user_repository.IsOtpLocked(lastOTP) {
		// 		return "too many attempts", nil, 403, errors.New("otp locked")
		// 	}
		//     ok, wait := utils.CanResendOTP(lastOTP.UpdatedAt)
		//     if !ok {
		// 	    return utils.OTPWaitError(wait), nil, 405, errors.New("otp wait error")
		//     }
		// }

		// otp := utils.GenerateOTP()
		// user_repository.UpsertOTP(s.DB, req.Email, otp)

		// err := email.SendEmail(otp, user).Verification()
		// if err != nil {
		// 	return "failed to send email", nil, http.StatusInternalServerError, err
		// }
		// user_repository.IncrementOtpHit(s.DB, req.Email)
	}

	// }
	return "OTP Sent Successfully", nil, http.StatusOK, nil
}


func HandleOTP(db *gorm.DB, identifier string, sendOTP OTPSendFunc,) (string, interface{}, int, error) {

	lastOTP, _ := user_repository.GetVerification(db, identifier)
	if lastOTP != nil {

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