package cryptography

import "time"

type Encrypting interface {
	JwtEncrypt(payload map[string]any, expirationAt time.Time, key string) (string, error)
	JwtDecrypt(tokenString string, key string) (map[string]any, error)
}
