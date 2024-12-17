package security

import (
	"crypto/rand"
	"encoding/hex"
)

const DefaultTokenLength = 32

func GenerateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
