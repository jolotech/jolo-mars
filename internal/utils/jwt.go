
package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email  string `json:"email"`
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type AdminClaims struct {
	AdminID uint   `json:"admin_id"`
	Email   string `json:"email"`
	Purpose string `json:"purpose"` // "access" or "pwd_change"
	jwt.RegisteredClaims
}

func GenerateAuthToken(email string, userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	expiry, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		expiry = 24 * time.Hour
	}

	claims := jwt.MapClaims{
		"email": email,
		"user_id": userID,
		"exp":   time.Now().Add(expiry).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func GenerateAdminAuthToken(email, purpose string, adminID uint) (string, error) {
	secret := os.Getenv("ADMIN_JWT_SECRET")

	expiry, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		expiry = 24 * time.Hour
	}

	claims := jwt.MapClaims{
		"email": email,
		"admin_id": adminID,
		"purpose": purpose,
		"exp":   time.Now().Add(expiry).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func SignAdminToken(secret string, adminID uint, email string, purpose string, ttl time.Duration) (string, error) {
	now := time.Now()
	
	claims := AdminClaims{
		AdminID: adminID,
		Email:   email,
		Purpose: purpose,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

