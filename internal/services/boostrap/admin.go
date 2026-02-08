package bootstrap_service

import (
	// "log"
	"os"
	"strings"

	// "time"

	"github.com/jolotech/jolo-mars/internal/models"
	"github.com/jolotech/jolo-mars/internal/repository/boostrap"
	admin_repo "github.com/jolotech/jolo-mars/internal/repository/admin"
	"github.com/jolotech/jolo-mars/internal/utils"
)

type BootstrapService struct {
	adminRepo *admin_repo.AdminAthRepo
	adminBoostrapRepo *admin_repository.AdminBoostrap
}



func NewBootstrapService(adminRepo *admin_repo.AdminAthRepo, adminBoostrapRepo *admin_repository.AdminBoostrap) *BootstrapService {
	return &BootstrapService{
		adminBoostrapRepo: adminBoostrapRepo,
		adminRepo: adminRepo,
	}
}

type BootstrapResult struct {
	Created      bool
	TempPassword string
	Reason       string
}

// func (s *BootstrapService) EnsureSuperAdminFromEnvSilently() (*BootstrapResult, error) {
// 	// Gate: only run when explicitly enabled
// 	log.Panicln("Check fail Silently 1")
// 	if strings.ToLower(os.Getenv("BOOTSTRAP_SUPER_ADMIN")) != "true" {
// 		log.Panicln("Check fail Silently 2 == false")

// 		return &BootstrapResult{Created: false}, nil
// 	}

// 	// Silent exit if any admin exists already (your request)
// 	exists, err := s.adminRepo.AnyAdminExists()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if exists {
// 		return &BootstrapResult{Created: false}, nil
// 	}

// 	name := strings.TrimSpace(os.Getenv("SUPER_ADMIN_NAME"))
// 	email := strings.TrimSpace(strings.ToLower(os.Getenv("SUPER_ADMIN_EMAIL")))
// 	pass := os.Getenv("SUPER_ADMIN_PASSWORD") // optional

// 	// If required env missing: fail silently as requested
// 	if name == "" || email == "" {
// 		return &BootstrapResult{Created: false}, nil
// 	}

// 	// If password not provided, generate temp password
// 	tempGenerated := ""
// 	if strings.TrimSpace(pass) == "" {
// 		tempGenerated = utils.GenerateStrongPassword(16)
// 		pass = tempGenerated
// 	}

// 	hash, err := utils.HashPassword(pass)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// now := time.Now()

// 	admin := &models.Admin{
// 		Name:               name,
// 		Email:              email,
// 		Password:           hash,
// 		Role:               "super-admin",
// 		MustChangePassword: true,
// 		PasswordChangedAt:  nil,
// 		TwoFAEnabled:       false,
// 		TwoFAConfirmedAt:   nil,
// 		LastLoginAt:        nil,
// 		// _                  : now,
// 	}

// 	if err := s.adminRepo.Create(admin); err != nil {
// 		return nil, err
// 	}

// 	return &BootstrapResult{
// 		Created:      true,
// 		TempPassword: tempGenerated,
// 	}, nil
// }


// func (s *BootstrapService) EnsureSuperAdminFromEnvSilently() (*BootstrapResult, error) {
// 	gate := strings.ToLower(strings.TrimSpace(os.Getenv("BOOTSTRAP_SUPER_ADMIN")))
// 	log.Println("Check fail Silently 1")

// 	if gate != "true" {
// 	    log.Println("Check fail Silently 2 == false")
// 		return &BootstrapResult{Created: false, Reason: "BOOTSTRAP_SUPER_ADMIN is not true"}, nil
// 	}

// 	exists, err := s.adminRepo.AnyAdminExists()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if exists {
// 		return &BootstrapResult{Created: false, Reason: "admin already exists"}, nil
// 	}

// 	name := strings.TrimSpace(os.Getenv("SUPER_ADMIN_NAME"))
// 	email := strings.TrimSpace(strings.ToLower(os.Getenv("SUPER_ADMIN_EMAIL")))
// 	pass := os.Getenv("SUPER_ADMIN_PASSWORD")

// 	if name == "" || email == "" {
// 		return &BootstrapResult{Created: false, Reason: "missing SUPER_ADMIN_NAME or SUPER_ADMIN_EMAIL"}, nil
// 	}

// 	temp := ""
// 	if strings.TrimSpace(pass) == "" {
// 		temp = utils.GenerateStrongPassword(16)
// 		pass = temp
// 	}

// 	hash, err := utils.HashPassword(pass)
// 	if err != nil {
// 		return nil, err
// 	}

// 	admin := &models.Admin{
// 		Name:               name,
// 		Email:              email,
// 		Password:           hash,
// 		Role:               "super-admin",
// 		MustChangePassword: true,
// 		TwoFAEnabled:       false,
// 	}

// 	if err := s.adminRepo.Create(admin); err != nil {
// 		return nil, err
// 	}

// 	return &BootstrapResult{Created: true, TempPassword: temp, Reason: "created"}, nil
// }


func (s *BootstrapService) EnsureSuperAdminFromEnvSilently() (*BootstrapResult, error) {
	gate := strings.ToLower(strings.TrimSpace(os.Getenv("BOOTSTRAP_SUPER_ADMIN")))
	if gate != "true" {
		return &BootstrapResult{Created: false, Reason: "BOOTSTRAP_SUPER_ADMIN is not true"}, nil
	}

	// ✅ Only block if SUPER ADMIN exists (not any admin)
	superExists, err := s.adminBoostrapRepo.AnySuperAdminExists()
	if err != nil {
		return nil, err
	}
	if superExists {
		return &BootstrapResult{Created: false, Reason: "super admin already exists"}, nil
	}

	name := strings.TrimSpace(os.Getenv("SUPER_ADMIN_NAME"))
	email := strings.TrimSpace(strings.ToLower(os.Getenv("SUPER_ADMIN_EMAIL")))
	pass := os.Getenv("SUPER_ADMIN_PASSWORD")

	if name == "" || email == "" {
		return &BootstrapResult{Created: false, Reason: "missing SUPER_ADMIN_NAME or SUPER_ADMIN_EMAIL"}, nil
	}

	// ✅ avoid duplicate email, even if role is manager/support
	emailExists, err := s.adminRepo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return &BootstrapResult{Created: false, Reason: "email already exists (cannot bootstrap super admin)"}, nil
	}

	temp := ""
	if strings.TrimSpace(pass) == "" {
		temp = utils.GenerateStrongPassword(16)
		pass = temp
	}

	hash, err := utils.HashPassword(pass)
	if err != nil {
		return nil, err
	}

	admin := &models.Admin{
		Name:               name,
		Email:              email,
		Password:           hash,
		Role:               "super-admin",
		MustChangePassword: true,
		TwoFAEnabled:       false,
	}

	if err := s.adminBoostrapRepo.CreateSuperAdmin(admin); err != nil {
		return nil, err
	}

	return &BootstrapResult{Created: true, TempPassword: temp, Reason: "created"}, nil
}
