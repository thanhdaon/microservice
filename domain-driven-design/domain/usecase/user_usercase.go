package usecase

import (
	"domain-driven-design/domain/repository"
	"errors"
)

type UserUsecase interface {
	Login(email, password string) (string, error)
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

type userUsecase struct {
	UserRepo repository.UserRepository
}

func (uc *userUsecase) Login(email, passwords string) (string, error) {
	user := uc.UserRepo.GetByEmail(email)
	if user == nil {
		return "", errors.New("user not found")
	}
	return "", nil
}
