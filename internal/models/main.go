package models

import (
	"crypto/rand"
	"encoding/base64"
)

// GeneratePublicID returns a URL-safe random string of exactly 15 characters
func GeneratePublicID() string {
	b := make([]byte, 11) // 11 bytes ≈ 15 chars base64
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // extremely rare; acceptable for ID generation
	}

	id := base64.RawURLEncoding.EncodeToString(b)
	return id[:15]
}