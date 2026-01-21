
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
