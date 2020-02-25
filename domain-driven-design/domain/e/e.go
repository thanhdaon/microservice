package e

import "errors"

var (
	USER_NOT_FOUND       = errors.New("user not found")
	WRONG_PASSWORD       = errors.New("wrong password")
	EMAIL_ALREADY_EXISTS = errors.New("email already exists")
)
