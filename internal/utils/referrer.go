package utils

import (
	"crypto/rand"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const referralCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// PUBLIC FUNCTION — call this to generate a unique referral code
func GenerateRefererCode(db *gorm.DB) (string) {
	for {
		code, err := generateUniqueReferralCode()
		if err != nil {
			return ""
		}

		if !referralCodeExists(db, code) {
			return code
		}
	}
}

func generateUniqueReferralCode() (string, error) {
	randomPart, err := randomString(10) 
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("RPC-%s-%s-%s",
		randomPart[0:4],
		randomPart[4:8],
		randomPart[8:10],
	), nil
}

func randomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] = referralCharset[int(b[i])%len(referralCharset)]
	}

	return strings.ToUpper(string(b)), nil
}

func referralCodeExists(db *gorm.DB, code string) bool {
	var count int64
	db.Table("users").
		Where("referral_code = ?", code).
		Count(&count)

	return count > 0
}
