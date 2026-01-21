package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAuthToken(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	expiry, _ := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))

	claims := AuthClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
