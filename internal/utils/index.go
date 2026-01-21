package utils

import (
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

func ValidateRegister(req RegisterRequest, db *gorm.DB) error {
	// Trim spaces
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Phone = strings.TrimSpace(req.Phone)

	// Basic required checks (extra safety)
	if req.Name == "" {
		return errors.New("name is required")
	}

	if req.Phone == "" {
		return errors.New("phone number is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	// Email validation (if provided)
	if req.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(req.Email) {
			return errors.New("invalid email address")
		}
	}

	// Check if email already exists
	if req.Email != "" {
		var count int64
		if err := db.Model(&User{}).
			Where("email = ?", req.Email).
			Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return errors.New("email already registered")
		}
	}

	// Check if phone already exists
	var phoneCount int64
	if err := db.Model(&User{}).
		Where("phone = ?", req.Phone).
		Count(&phoneCount).Error; err != nil {
		return err
	}

	if phoneCount > 0 {
		return errors.New("phone number already registered")
	}

	return nil
}
