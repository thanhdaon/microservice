package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (repo *repository) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
	`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	if _, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string

	if err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email); err != nil {
		return "", RepoErr
	}

	return email, nil
}
