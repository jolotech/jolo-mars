package bootstrap_service

import (
	"os"
	"strings"
	// "time"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/utils"
)

type BootstrapService struct {
	adminRepo *admin_repository.AdminRepo
}

func NewBootstrapService(adminRepo *admin_repository.AdminRepo) *BootstrapService {
	return &BootstrapService{adminRepo: adminRepo}
}

type BootstrapResult struct {
	Created      bool
	TempPassword string 
}

func (s *BootstrapService) EnsureSuperAdminFromEnvSilently() (*BootstrapResult, error) {
	// Gate: only run when explicitly enabled
	if strings.ToLower(os.Getenv("BOOTSTRAP_SUPER_ADMIN")) != "true" {
		return &BootstrapResult{Created: false}, nil
	}

	// Silent exit if any admin exists already (your request)
	exists, err := s.adminRepo.AnyAdminExists()
	if err != nil {
		return nil, err
	}
	if exists {
		return &BootstrapResult{Created: false}, nil
	}

	name := strings.TrimSpace(os.Getenv("SUPER_ADMIN_NAME"))
	email := strings.TrimSpace(strings.ToLower(os.Getenv("SUPER_ADMIN_EMAIL")))
	pass := os.Getenv("SUPER_ADMIN_PASSWORD") // optional

	// If required env missing: fail silently as requested
	if name == "" || email == "" {
		return &BootstrapResult{Created: false}, nil
	}

	// If password not provided, generate temp password
	tempGenerated := ""
	if strings.TrimSpace(pass) == "" {
		tempGenerated = utils.GenerateStrongPassword(16)
		pass = tempGenerated
	}

	hash, err := utils.HashPassword(pass)
	if err != nil {
		return nil, err
	}

	// now := time.Now()

	admin := &models.Admin{
		Name:               name,
		Email:              email,
		Password:           hash,
		Role:               "super-admin",
		MustChangePassword: true,
		PasswordChangedAt:  nil,
		TwoFAEnabled:       false,
		TwoFAConfirmedAt:   nil,
		LastLoginAt:        nil,
		// _                  : now,
	}

	if err := s.adminRepo.Create(admin); err != nil {
		return nil, err
	}

	return &BootstrapResult{
		Created:      true,
		TempPassword: tempGenerated,
	}, nil
}
