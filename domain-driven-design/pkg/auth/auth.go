package auth

import (
	"domain-driven-design/domain/helper"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthHelper() helper.Auth {
	return &authHelper{}
}

type authHelper struct{}

func (authHelper) CreateToken(userid uint64) (string, error) {
	return fmt.Sprintf("token %d", userid), nil
}

func (authHelper) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (authHelper) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
