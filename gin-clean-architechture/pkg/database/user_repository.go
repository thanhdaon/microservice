package database

import (
	"domain-driven-design/domain/e"
	"domain-driven-design/domain/entity"
	"domain-driven-design/domain/repository"
	"strings"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Save(user *entity.User) (*entity.User, error) {
	var err error
	if user.ID == 0 {
		err = r.db.Create(user).Error
	} else {
		err = r.db.Save(user).Error
	}

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, e.EMAIL_ALREADY_EXISTS
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetByID(id uint64) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, "email = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepo) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, e.USER_NOT_FOUND
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
