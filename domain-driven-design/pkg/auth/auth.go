package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	CreateToken(uint64) (string, error)
	HashPassword(string) ([]byte, error)
	VerifyPassword(string, string) error
}

func NewAuthUtil() AuthUtil {
	return &authUtil{}
}

type authUtil struct{}

func (authUtil) CreateToken(userid uint64) (string, error) {
	return fmt.Sprintf("token %d", userid), nil
}

func (authUtil) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (authUtil) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
