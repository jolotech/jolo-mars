package utils

import (
	"fmt"
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

func GenerateUserAuthToken(email string, userID uint) (string, error) {
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

	expiry := 24 * time.Hour

	switch purpose {
	case "pwd_change":
		expiry = 15 * time.Minute
		if cfg.AuthPassExpIn != "" {
			if d, err := time.ParseDuration(cfg.AuthPassExpIn); err == nil {
				expiry = d
			}
		}

	case "access":
		expiry = 24 * time.Hour
		if cfg.AuthExpIn != "" {
			if d, err := time.ParseDuration(cfg.AuthExpIn); err == nil {
				expiry = d
			}
		}

	default:
		return "", fmt.Errorf("invalid token purpose: %s", purpose)
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"email":    email,
		"admin_id": adminID,
		"purpose":  purpose,
		"exp":      now.Add(expiry).Unix(),
		"iat":      now.Unix(),
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

