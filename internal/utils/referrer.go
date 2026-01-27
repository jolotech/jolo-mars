package utils

import (
	"crypto/rand"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const referralCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// PUBLIC FUNCTION — call this
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

// INTERNAL — builds the code
func generateUniqueReferralCode() (string, error) {
	randomPart, err := randomString(10) // ensures total length ≥ 12
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("RPCJ%sO%sL%sO",
		randomPart[0:4],
		randomPart[4:8],
		randomPart[8:10],
	), nil
}

// INTERNAL — crypto-safe random
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
// INTERNAL — checks DB for existing code

func referralCodeExists(db *gorm.DB, code string) bool {
	var count int64
	db.Table("users").
		Where("referral_code = ?", code).
		Count(&count)

	return count > 0
}
