package helpers

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/maps"

	"github.com/my-storage/ms-profile/src/shared/protocols/cryptography"
	utilsTime "github.com/my-storage/ms-profile/src/shared/utils/time"
)

type JwtAdapter struct {
	cryptography.Encrypting
}

func (j *JwtAdapter) JwtEncrypt(payload map[string]any, expirationAt time.Time, key string) (string, error) {
	claims := jwt.MapClaims{
		"alg": "HS256",
		"typ": "JWT",
		"exp": utilsTime.TimeToMilliseconds(expirationAt),
	}

	maps.Copy(claims, payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func (j *JwtAdapter) JwtDecrypt(tokenString string, key string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing model: %v", token.Header["alg"])
		}

		return []byte(key), nil
	}, jwt.WithValidMethods([]string{"HMAC-SHA256"}))

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
