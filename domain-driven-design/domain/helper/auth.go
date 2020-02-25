package helper

import "github.com/dgrijalva/jwt-go"

type Auth interface {
	CreateToken(string) (string, error)
	ParseToken(string) (*Claims, error)
	HashPassword(string) ([]byte, error)
	VerifyPassword(string, string) error
}

type Claims struct {
	Email string
	jwt.StandardClaims
}
