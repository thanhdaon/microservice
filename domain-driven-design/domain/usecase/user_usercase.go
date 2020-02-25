package usecase

import (
	"domain-driven-design/domain/repository"
	"domain-driven-design/utils/auth"
	"errors"
)

type UserUsecase interface {
	Login(email, password string) (string, error)
}

func NewUserUsecase(authUtil auth.AuthUtil, userRepo repository.UserRepository) UserUserUsecase {
	return &userUsecase{authUtil, userRepo}
}

type userUsecase struct {
	AuthUtil auth.AuthUtil
	UserRepo repository.UserRepository
}

func (uc *userUsecase) Login(email, passwords string) (string, error) {
	user := uc.UserRepo.GetByEmail(email)
	if user == nil {
		return "", errors.New("user not found")
	}

}
