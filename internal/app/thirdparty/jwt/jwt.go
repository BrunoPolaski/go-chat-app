package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct{}

func (ja *JWTAdapter) GenerateToken(payload jwt.Claims) (string, error) {
	k := os.Getenv("JWT_SECRET")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := t.SignedString([]byte(k))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (ja *JWTAdapter) ParseToken() {
	// TODO: Implement this method
}
