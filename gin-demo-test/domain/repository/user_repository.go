package repository

import (
	"errors"
	"gin-demo-test/domain/entity"
	"gin-demo-test/utils/security"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("email is already taken!")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *userRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (r *userRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("password is not correct")
	}
	return &user, nil
}
