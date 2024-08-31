package jwt

import "github.com/golang-jwt/jwt/v5"

type JWT interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ParseToken(tokenString string, claims jwt.Claims) error
}
