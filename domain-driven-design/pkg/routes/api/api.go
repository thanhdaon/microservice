package api

import (
	"domain-driven-design/domain/usecase"
)

var userUC usecase.UserUsecase

type Dependences struct {
	UserUC usecase.UserUsecase
}

func Setup(dependences Dependences) {
	userUC = dependences.UserUC
}
