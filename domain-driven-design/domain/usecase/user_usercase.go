package usecase

import (
	"domain-driven-design/domain/e"
	"domain-driven-design/domain/helper"
	"domain-driven-design/domain/repository"
	"fmt"
)

type UserUsecase interface {
	Login(email, password string) (string, error)
}

func NewUserUsecase(userRepo repository.UserRepository, authHelper helper.Auth) UserUsecase {
	return &userUsecase{userRepo, authHelper}
}

type userUsecase struct {
	UserRepo   repository.UserRepository
	AuthHelper helper.Auth
}

func (uc *userUsecase) Login(email, passwords string) (string, error) {
	user, err := uc.UserRepo.GetByEmail(email)
	switch err {
	case nil:
		return fmt.Sprintf("token-%s-%s", user.FirstName, user.LastName), nil
	case e.USER_NOT_FOUND:
		return "", e.USER_NOT_FOUND
	default:
		return "", err
	}
}
