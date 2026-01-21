package helpers


import (
	// "crypto/rsa"
	// "encoding/json"
	"time"

	"github.com/jolotech/jolo-mars/types"
	"github.com/golang-jwt/jwt/v5"
)

func createJWT(sa *types.FirebaseServiceAccount) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(sa.PrivateKey))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"iss":   sa.ClientEmail,
		"scope": "https://www.googleapis.com/auth/firebase.messaging",
		"aud":   "https://oauth2.googleapis.com/token",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
