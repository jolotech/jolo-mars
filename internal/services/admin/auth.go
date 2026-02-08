package admin_services

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/jolotech/jolo-mars/config"
	"github.com/jolotech/jolo-mars/internal/repository/admin"
	// "github.com/jolotech/jolo-mars/internal/utils/jwt"
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
		setupToken, err := utils.GenerateAdminAuthToken(admin.Email, "pwd_change", admin.ID, 15*time.Minute)
		if err != nil {
			return "failed to create token", nil, http.StatusInternalServerError, err
		}

		return "password change required", types.AdminLoginResponse{
			PasswordChangeRequired: true,
			SetupToken:             setupToken,
		}, http.StatusOK, nil
	}

	// Normal access token
	accessToken, err := jwt.SignAdminToken(cfg.JWTSecret, admin.ID, admin.Email, "access", 24*time.Hour)
	if err != nil {
		return "failed to create token", nil, http.StatusInternalServerError, err
	}

	return "login successful", types.AdminLoginResponse{
		AccessToken:            accessToken,
		PasswordChangeRequired: false,
	}, http.StatusOK, nil
}

// Uses SetupToken from Authorization: Bearer <token>
func (s *AdminAuthService) ChangePassword(ctx context.Context, setupToken string, req types.AdminChangePasswordRequest) (string, any, int, error) {
	cfg := config.LoadConfig()

	claims, err := jwt.ParseAdminToken(cfg.JWTSecret, setupToken)
	if err != nil {
		return "invalid token", nil, http.StatusUnauthorized, err
	}
	if claims.Purpose != "pwd_change" {
		return "invalid token purpose", nil, http.StatusUnauthorized, errors.New("invalid token purpose")
	}

	admin, err := s.repo.GetByEmail(strings.ToLower(strings.TrimSpace(claims.Email)))
	if err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}
	if admin == nil || admin.ID != claims.AdminID {
		return "invalid token", nil, http.StatusUnauthorized, errors.New("invalid token")
	}

	// Verify current password
	if !security.VerifyPassword(admin.Password, req.CurrentPassword) {
		return "current password incorrect", nil, http.StatusUnauthorized, errors.New("wrong password")
	}

	// Strong password check
	ok, msg := security.IsStrongPassword(req.NewPassword)
	if !ok {
		return msg, nil, http.StatusBadRequest, errors.New(msg)
	}

	// Hash and update
	hash, err := security.HashPassword(req.NewPassword)
	if err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}

	if err := s.repo.UpdatePassword(admin.ID, hash, false); err != nil {
		return "failed", nil, http.StatusInternalServerError, err
	}

	// Optionally issue access token immediately after change
	accessToken, err := jwt.SignAdminToken(cfg.JWTSecret, admin.ID, admin.Email, "access", 24*time.Hour)
	if err != nil {
		return "password updated but token failed", nil, http.StatusInternalServerError, err
	}

	return "password updated successfully", map[string]any{
		"access_token": accessToken,
	}, http.StatusOK, nil
}
