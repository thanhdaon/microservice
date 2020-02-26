package usecase

import (
	"domain-driven-design/domain/e"
	"domain-driven-design/domain/entity"
	"domain-driven-design/domain/helper"
	"domain-driven-design/domain/repository"
)

type AuthUsecase interface {
	Signin(email, password string) (token string, err error)
	Signup(email, password, fistname, lastname string) (*entity.User, error)
}

func NewUserUsecase(userRepo repository.UserRepository, authHelper helper.Auth) AuthUsecase {
	return &authUsecase{userRepo, authHelper}
}

type authUsecase struct {
	UserRepo   repository.UserRepository
	AuthHelper helper.Auth
}

func (uc *authUsecase) Signin(email, password string) (string, error) {
	user, err := uc.UserRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if err := uc.AuthHelper.VerifyPassword(user.Password, password); err != nil {
		return "", e.WRONG_PASSWORD
	}

	token, err := uc.AuthHelper.CreateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *authUsecase) Signup(email, password, firstname, lastname string) (*entity.User, error) {
	hashPassword, err := uc.AuthHelper.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := uc.UserRepo.Save(&entity.User{
		Email:     email,
		Password:  string(hashPassword),
		FirstName: firstname,
		LastName:  lastname,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
