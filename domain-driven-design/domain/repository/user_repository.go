package repository

import "domain-driven-design/domain/entity"

type UserRepository interface {
	Save(*entity.User) *entity.User
	GetAll() []entity.User
	GetByID(uint64) *entity.User
	GetByEmail(string) *entity.User
}
