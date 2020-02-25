package auth

import (
	"domain-driven-design/domain/helper"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthHelper(jwtSecret []byte) helper.Auth {
	return &authHelper{jwtSecret}
}

type authHelper struct {
	JwtSecret []byte
}

func (a *authHelper) CreateToken(email string) (string, error) {
	claims := helper.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(a.JwtSecret)
}

func (a *authHelper) ParseToken(token string) (*helper.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&helper.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return a.JwtSecret, nil
		},
	)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*helper.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (authHelper) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (authHelper) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
