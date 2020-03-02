package auth

import (
	"domain-driven-design/domain/helper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthHelper(t *testing.T) {
	auth := NewAuthHelper("jwt_secret")
	assert.NotNil(t, auth)
}

func TestJWT_Success(t *testing.T) {
	auth := authHelper{[]byte("jwt_secret")}

	originString := "hello_world"
	token, err := auth.CreateToken(originString, 3*time.Hour)
	assert.Nil(t, err)

	claims, err := auth.ParseToken(token)
	assert.Nil(t, err)

	assert.IsType(t, &helper.Claims{}, claims)
	assert.Equal(t, originString, claims.Email)
}

func TestJWTParseToken_Fail(t *testing.T) {
	auth := authHelper{[]byte("jwt_secret")}
	claims, err := auth.ParseToken("okok")
	assert.Nil(t, claims)
	assert.NotNil(t, err)
}

func TestPasswordHashing(t *testing.T) {
	auth := authHelper{[]byte("jwt_secret")}

	originString := "hello_world"

	hashed, err := auth.HashPassword(originString)
	assert.Nil(t, err)

	err = auth.VerifyPassword(hashed, originString)
	assert.Nil(t, err)

	err = auth.VerifyPassword(hashed, "okok")
	assert.NotNil(t, err)
}
