package usecase

import (
	"domain-driven-design/domain/helper"
	"domain-driven-design/domain/repository"
	"domain-driven-design/pkg/e"
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
	user := uc.UserRepo.GetByEmail(email)
	if user == nil {
		return "", e.USER_NOT_FOUND
	}
	return fmt.Sprintf("token-%s-%s", user.FirstName, user.LastName), nil
}
