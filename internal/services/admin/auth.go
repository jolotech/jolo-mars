package admin_services

import (
	"errors"
	"net/http"
	"strings"
	"encoding/base64"
	"os"

	// "github.com/jolotech/jolo-mars/internal/repository"
	// "github.com/jolotech/jolo-mars/internal/utils"
	// "github.com/jolotech/jolo-mars/types"

	"github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/utils"
	"github.com/jolotech/jolo-mars/types"
)

type AdminAuthService struct {
	adminAuthRepo *admin_repository.AdminAuthRepo
}

func NewAdminAuthService(adminAuthRepo *admin_repository.AdminAuthRepo) *AdminAuthService {
	return &AdminAuthService{adminAuthRepo: adminAuthRepo}
}



func (s *AdminAuthService) Login(req types.AdminLoginRequest) (string, any, int, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	admin, err := s.adminAuthRepo.GetByEmail(email)
	if err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}
	if admin == nil {
		return "invalid credentials", nil, http.StatusUnauthorized, errors.New("invalid credentials")
	}

	if !utils.ComparePassword(admin.Password, req.Password) {
		return "invalid credentials", nil, http.StatusUnauthorized, errors.New("invalid credentials")
	}

	setupToken, err := utils.GenerateAdminAuthToken(admin.Email, "2FA", admin.PublicID)

	data := types.AdminLoginResponse{
		    Requires2FA: true,
	        Requires2FAMessage: "2FA is required for this account",
		    TwoFAToken:  setupToken,
	    }

	if admin.TwoFAEnabled {
        if err != nil {
			return "failed to create token", nil, http.StatusInternalServerError, err
		}
	    return "2FA required", data, 200, nil
	}

	data.Requires2FAMessage = "2FA not setup for this account, please setup 2FA to secure your account"

	if !admin.TwoFAEnabled{
		return "2FA not setup Please use the setup 2fa endpoint", data, http.StatusForbidden, errors.New("2FA not setup")
	}

	// Normal access token
	return "account malfunction", nil, http.StatusBadRequest, errors.New("bad or Admin account malfunction")
}

func (s *AdminAuthService) Setup2FA(adminId string) (string, any, int, error) {
	admin, err := s.adminAuthRepo.GetByPublicID(adminId)
	if err != nil || admin == nil {
		return "failed", types.AdminTwoFASetupResponse{}, http.StatusInternalServerError, err
	}

	key, err := utils.Generate2faTOTPKey("Jolo Admin", admin.Email)
	if err != nil {
		return "failed", types.AdminTwoFASetupResponse{}, http.StatusInternalServerError, err
	}

	encKeyB64 := os.Getenv("TWO_FA_ENC_KEY")
	encKey, err := base64.StdEncoding.DecodeString(encKeyB64)
	if err != nil || len(encKey) != 32 {
		return "failed", nil, http.StatusInternalServerError, errors.New("invalid TWO_FA_ENC_KEY (must be base64 of 32 bytes)")
	}

	encSecret, err := utils.EncryptString(key.Secret(), encKey)
	if err != nil {
		return "failed", types.AdminTwoFASetupResponse{}, http.StatusInternalServerError, err
	}

	if err := s.adminAuthRepo.Save2FASecret(admin.ID, encSecret); err != nil {
		return "failed", types.AdminTwoFASetupResponse{}, http.StatusInternalServerError, err
	}

	return "2fa setup successful", types.AdminTwoFASetupResponse{OtpAuthURL: key.URL()}, http.StatusOK, nil
}


func (s *AdminAuthService) Confirm2FA(adminId, code string) (string, any, int, error) {
	admin, err := s.adminAuthRepo.GetByPublicID(adminId)
	if err != nil || admin == nil{
		return "failed", nil, http.StatusInternalServerError, err
	}

	if admin.TwoFASecretEnc == "" {
		return "2fa not initialized", nil, http.StatusForbidden, errors.New("2fa not initialized")
	}

	encKeyB64 := os.Getenv("TWO_FA_ENC_KEY")
	encKey, _ := base64.StdEncoding.DecodeString(encKeyB64)

	secret, err := utils.DecryptString(admin.TwoFASecretEnc, encKey)
	if err != nil {
		return "failed to decrypt 2fa secret", nil, http.StatusInternalServerError, errors.New("failed to decrypt 2fa secret")
	}

	if !utils.Verify2faTOTP(code, secret) {
		return "invalid 2fa code", nil, http.StatusUnauthorized, errors.New("invalid 2fa code")
	}

	if err := s.adminAuthRepo.Enable2FA(admin.ID); err != nil {
		return "failed to enable 2fa", nil, http.StatusInternalServerError, err
	}

	// Must change password first
	if admin.MustChangePassword {
		setupToken, err := utils.GenerateAdminAuthToken(admin.Email, "pwd_change", admin.PublicID)
		if err != nil {
			return "failed to create token", nil, http.StatusInternalServerError, err
		}

		data := types.AdminLoginResponse{
			PasswordChangeRequired: true,
			SetupToken:             setupToken,
		}

		return "password change required", data, http.StatusOK, nil
	}

	// Normal access token
	accessToken, err := utils.GenerateAdminAuthToken(admin.Email, "access", admin.PublicID)
	if err != nil {
		return "failed to create token", nil, http.StatusInternalServerError, err
	}

	data := types.AdminLoginResponse{
		AccessToken:            accessToken,
		PasswordChangeRequired: false,
		Admin: admin,
	}

	return "login successful", data, http.StatusOK, nil
}

func (s *AdminAuthService) ChangePassword(req types.AdminChangePasswordRequest) (string, any, int, error) {

	admin, err := s.adminAuthRepo.GetByEmail(req.Email)
	if err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}

	if req.NewPassword != req.ConfirmPassword {
		return "password does not match", nil, http.StatusBadRequest, errors.New("pasword mismatch")
	}

	// Verify current password
	if !utils.ComparePassword(admin.Password, req.CurrentPassword) {
		return "current password incorrect", nil, http.StatusUnauthorized, errors.New("wrong password")
	}

	// Strong password check
	ok, msg := utils.IsStrongPassword(req.NewPassword)
	if !ok {
		return msg, nil, http.StatusBadRequest, errors.New(msg)
	}

	// Hash and update
	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}

	if err := s.adminAuthRepo.UpdatePassword(admin.ID, hash, false); err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}

	// issue access token immediately after change
	accessToken, err := utils.GenerateAdminAuthToken(admin.Email, "access", admin.PublicID)
	if err != nil {
		return "password updated but access granted failed", nil, http.StatusInternalServerError, err
	}

	data := types.AdminLoginResponse{
		AccessToken: accessToken,
		Admin: admin,
	}
	return "password updated successfully", data, http.StatusOK, nil
}

func (s *AdminAuthService) DeleteAdminByEmail(email string) (string, any, int, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	admin, err := s.adminAuthRepo.DeleteByEmail(email)

	if err != nil || admin == nil {
		return "failed to delete admin", nil, http.StatusInternalServerError, err
	}
	return "admin deleted successfully", admin, http.StatusOK, nil
}