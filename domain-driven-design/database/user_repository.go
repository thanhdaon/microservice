package database

import (
	"domain-driven-design/domain/entity"
	"domain-driven-design/domain/repository"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Save(user *entity.User) *entity.User {
	if user.ID == 0 {
		r.db.Create(user)
	} else {
		r.db.Save(user)
	}
	return user
}

func (r *userRepo) GetByID(id uint64) *entity.User {
	var user entity.User
	if r.db.First(&user, "id = ?", id).RecordNotFound() {
		return nil
	}
	return &user
}

func (r *userRepo) GetAll() []entity.User {
	var users []entity.User
	r.db.Find(&users)
	return users
}

func (r *userRepo) GetByEmail(email string) *entity.User {
	var user entity.User
	if r.db.Where("email = ?", email).Take(&user).RecordNotFound() {
		return nil
	}
	return &user
}
