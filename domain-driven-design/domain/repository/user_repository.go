package repository

import "domain-driven-design/domain/entity"

type UserRepository interface {
	Save(*entity.User) (*entity.User, error)
	GetAll() ([]entity.User, error)
	GetByID(uint64) (*entity.User, error)
	GetByEmail(string) (*entity.User, error)
}
