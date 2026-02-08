package utils

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"math/big"
	"regexp"
)

// Hash password (SIGNUP)
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // cost = 10
	)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(hashedPassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(rawPassword),
	)
	return err == nil
}



func GenerateStrongPassword(length int) string {
	const chars = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz23456789!@#$%&*"
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		out[i] = chars[n.Int64()]
	}
	return string(out)
}

// Strong password policy:
// - >= 10 chars
// - at least 1 upper, 1 lower, 1 number, 1 special
func IsStrongPassword(p string) (bool, string) {
	if len(p) < 10 {
		return false, "new password must be at least 10 characters"
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(p) {
		return false, "new password must include at least one lowercase letter"
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(p) {
		return false, "new password must include at least one uppercase letter"
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(p) {
		return false, "new password must include at least one number"
	}
	if !regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(p) {
		return false, "new password must include at least one special character"
	}
	return true, ""
}
