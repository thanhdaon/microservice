package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth interface {
	CreateToken(string, time.Duration) (string, error)
	ParseToken(string) (*Claims, error)
	HashPassword(string) (string, error)
	VerifyPassword(string, string) error
}

type Claims struct {
	Email string
	jwt.StandardClaims
}
