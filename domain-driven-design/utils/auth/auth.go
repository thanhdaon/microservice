package auth

import "fmt"

type AuthUtil interface {
	CreateToken(uint64) (string, error)
}

func NewAuthUtil() AuthUtil {
	return &authUtil{}
}

type authUtil struct{}

func (authUtil) CreateToken(userid uint64) (string, error) {
	return fmt.Sprintf("token %d", userid), nil

}
