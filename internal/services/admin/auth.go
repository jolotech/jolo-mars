package admin_services

import (
	"errors"
	"net/http"
	"strings"

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

	// Must change password first
	if admin.MustChangePassword {
		setupToken, err := utils.GenerateAdminAuthToken(admin.Email, "pwd_change", admin.ID)
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
	accessToken, err := utils.GenerateAdminAuthToken(admin.Email, "access", admin.ID)
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

// Uses SetupToken from Authorization: Bearer <token>
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
	accessToken, err := utils.GenerateAdminAuthToken(admin.Email, "access", admin.ID)
	if err != nil {
		return "password updated but access granted failed", nil, http.StatusInternalServerError, err
	}

	data := types.AdminLoginResponse{
		AccessToken: accessToken,
		Admin: admin,
	}
	return "password updated successfully", data, http.StatusOK, nil
}
