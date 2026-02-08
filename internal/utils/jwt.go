package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jolotech/jolo-mars/config"
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
	cfg := config.LoadConfig()

	expiry, err := time.ParseDuration(cfg.AuthExpIn)
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

	return token.SignedString([]byte(cfg.AuthSecret))
}

func GenerateAdminAuthToken(email, purpose string, adminID uint) (string, error) {
	cfg := config.LoadConfig()

	var expiry time.Duration

	if purpose == "pwd_change"{
		expiry, err := time.ParseDuration(cfg.AuthPassExpIn)
		if err != nil {
		    expiry = 24 * time.Hour
		}
	}
	if purpose == "auth_token" {
	    expiry, err := time.ParseDuration(cfg.AuthExpIn)
		if err != nil {
		    expiry = 15 * time.Hour
		}
	}


	claims := jwt.MapClaims{
		"email": email,
		"admin_id": adminID,
		"purpose": purpose,
		"exp":   time.Now().Add(expiry).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.AdminAuthSecret))
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

