package function

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewToken(id string, exp string, iss string, sub string, aud string, key string) string {

	now := time.Now()
	duration, _ := time.ParseDuration(exp + "m")
	expiresAt := now.Add(duration)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    iss,
		Subject:   sub,
		ID:        id,
		Audience:  []string{aud},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString([]byte(key))

	return accessToken
}
