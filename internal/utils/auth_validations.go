package utils

import (
	"errors"
	"regexp"
	"strings"
	"github.com/jolotech/jolo-mars/types"
	
	"gorm.io/gorm"
)




func ValidateRegister(req types.RegisterRequest, db *gorm.DB) string {
	// Trim spaces
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Phone = strings.TrimSpace(req.Phone)

	// Basic required checks (extra safety)
	if req.Name == "" {
		return "name is required"
	}

	if req.Phone == "" {
		return "phone number is required"
	}

	if len(req.Password) < 8 {
		return "password must be at least 8 characters"
	}

	// Email validation (if provided)
	if req.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(req.Email) {
			return "invalid email address"
		}
	}

	// Check if email already exists
	if req.Email != "" {
		var count int64
		if err := db.Table("users").
			Where("email = ?", req.Email).
			Count(&count).Error; err != nil {
			return "internal error"
		}

		if count > 0 {
			return "email already registered"
		}
	}

	// Check if phone already exists
	var phoneCount int64
	if err := db.Table("users").
		Where("phone = ?", req.Phone).
		Count(&phoneCount).Error; err != nil {
		return "internal error"
	}

	if phoneCount > 0 {
		return "phone number already registered"
	}

	return ""
}
