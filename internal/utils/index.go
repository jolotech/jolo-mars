package utils

import (
	"errors"
	"regexp"
	"strings"
	"github.com/jolotech/jolo-mars/types"
	
	"gorm.io/gorm"
)


func GenerateRefererCode(name string, id uint) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}

	// Generate a simple referral code using the name and ID
	nameParts := strings.Split(name, " ")
	firstName := nameParts[0]
	
	return firstName + "REF" + string(id), nil
}