package security

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/wolfsblu/go-chef/env"
	"time"
)

const CookieName = "SESSID"

func NewSessionCookie(payload string) string {
	expiry := 7 * 24 * 60 * time.Second // One week
	return fmt.Sprintf("%s=%s; HttpOnly; Secure; SameSite=strict; Path=/; Max-Age=%d", CookieName, payload, expiry)
}

func EncryptUserId(userId int64) (string, error) {
	var s = securecookie.New(
		[]byte(env.MustGet("COOKIE_HASH_KEY")),
		[]byte(env.MustGet("COOKIE_BLOCK_KEY")),
	)
	encoded, err := s.Encode(CookieName, userId)
	if err != nil {
		return "", err
	}
	return encoded, nil
}

func DecryptUserId(cookieValue string) (int64, error) {
	var userId int64
	var s = securecookie.New(
		[]byte(env.MustGet("COOKIE_HASH_KEY")),
		[]byte(env.MustGet("COOKIE_BLOCK_KEY")),
	)
	err := s.Decode(CookieName, cookieValue, &userId)
	if err != nil {
		return -1, err
	}
	return userId, nil
}
