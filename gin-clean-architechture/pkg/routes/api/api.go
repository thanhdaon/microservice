package api

import (
	"domain-driven-design/domain/usecase"
)

var authUC usecase.AuthUsecase

type response struct {
	Ok   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Dependences struct {
	AuthUC usecase.AuthUsecase
}

func Setup(dependences Dependences) {
	authUC = dependences.AuthUC
}
