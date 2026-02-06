package utils

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"math/big"
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
