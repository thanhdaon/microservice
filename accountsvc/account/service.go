package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	id := getID()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("error", err)
		return "Fail", err
	}

	logger.Log("create user", id)
	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("error", err)
	}

	logger.Log("Get user", id)

	return email, nil
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func getID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
