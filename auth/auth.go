package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

// Manager signs and verifies simple HMAC-based tokens.
// Token format: base64(payload).base64(hmacSHA256(payload, secret))
// This is a deliberately minimal example for educational purposes and does not
// implement expiration, rotation, or revocation.
type AuthManager interface {
	Sign(payload string) string
	Verify(token string) (string, error)
}
type AuthManagerImpl struct {
	secret []byte
}

// New creates a new Manager using the provided secret key.
func New(secret string) AuthManager {
	return &AuthManagerImpl{secret: []byte(secret)}
}

// Sign creates a token for the provided payload string.
func (m *AuthManagerImpl) Sign(payload string) string {
	sig := signHMAC([]byte(payload), m.secret)
	return base64.RawURLEncoding.EncodeToString([]byte(payload)) + "." + base64.RawURLEncoding.EncodeToString(sig)
}

// Verify checks a token and returns the original payload if valid.
func (m *AuthManagerImpl) Verify(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return "", errors.New("invalid token format")
	}
	payloadB, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return "", errors.New("invalid payload encoding")
	}
	sigB, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", errors.New("invalid signature encoding")
	}
	if !hmac.Equal(sigB, signHMAC(payloadB, m.secret)) {
		return "", errors.New("signature mismatch")
	}
	return string(payloadB), nil
}

func signHMAC(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
