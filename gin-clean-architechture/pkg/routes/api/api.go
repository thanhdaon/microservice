package api

import (
	"domain-driven-design/domain/usecase"
)

var userUC usecase.UserUsecase

type response struct {
	Ok   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Dependences struct {
	UserUC usecase.UserUsecase
}

func Setup(dependences Dependences) {
	userUC = dependences.UserUC
}
