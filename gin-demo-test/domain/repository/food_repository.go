package repository

import (
	"errors"
	"gin-demo-test/domain/entity"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
)

type FoodRepository interface {
	SaveFood(*entity.Food) (*entity.Food, map[string]string)
	GetFood(uint64) (*entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	UpdateFood(*entity.Food) (*entity.Food, map[string]string)
	DeleteFood(uint64) error
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepo{db}
}

type foodRepo struct {
	db *gorm.DB
}

func (r *foodRepo) SaveFood(food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}
	food.FoodImage = os.Getenv("DO_SPACES_URL") + food.FoodImage

	err := r.db.Debug().Create(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "food title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return food, nil
}

func (r *foodRepo) GetFood(id uint64) (*entity.Food, error) {
	var food entity.Food
	err := r.db.Debug().Where("id = ?", id).Take(&food).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("food not found")
	}
	return &food, nil
}

func (r *foodRepo) GetAllFood() ([]entity.Food, error) {
	var foods []entity.Food
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&foods).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return foods, nil
}

func (r *foodRepo) UpdateFood(food *entity.Food) (*entity.Food, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&food).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return food, nil
}

func (r *foodRepo) DeleteFood(id uint64) error {
	var food entity.Food
	err := r.db.Debug().Where("id = ?", id).Delete(&food).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
