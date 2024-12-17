package handlers

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/wolfsblu/go-chef/infra/env"
	"time"
)

const CookieName = "SESSID"

func createSessionCookie(userId int64) (string, error) {
	payload, err := encryptUserId(userId)
	if err != nil {
		return "", err
	}
	expiry := 7 * 24 * time.Hour // One week
	return fmt.Sprintf(
		"%s=%s; HttpOnly; Secure; SameSite=strict; Path=/; Max-Age=%d", CookieName, payload, int64(expiry/time.Second),
	), nil
}

func expireSessionCookie() string {
	return fmt.Sprintf(
		"%s=; HttpOnly; Secure; SameSite=strict; Path=/; Max-Age=%d", CookieName, 0,
	)
}

func getUserFromSessionCookie(cookieValue string) (int64, error) {
	return decryptUserId(cookieValue)
}

func encryptUserId(userId int64) (string, error) {
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

func decryptUserId(cookieValue string) (int64, error) {
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
