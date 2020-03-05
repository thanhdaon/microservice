package database

import (
	"domain-driven-design/domain/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func NewTestDBConnection() *gorm.DB, sqlmock.Sqlmock {
	db, mock, _ := sqlmock.New()
	gormdb, _ := gorm.Open("postgres", db)
	gormdb.SingularTable(true)
	gormdb.DropTableIfExists(&entity.User{})
	gormdb.AutoMigrate(&entity.User{})
	return gormdb, mock
}

func seedUser(db *gorm.DB) (*entity.User, error) {
	user := &entity.User{
		ID:        1,
		FirstName: "dao",
		LastName:  "thanh",
		Email:     "thanhdao@example.com",
		Password:  "aaa123",
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func seedUsers(db *gorm.DB) ([]entity.User, error) {
	users := []entity.User{
		{
			ID:        1,
			FirstName: "thanh",
			LastName:  "dao",
			Email:     "thanhdao@example.com",
			Password:  "aaa123",
		},
		{
			ID:        2,
			FirstName: "tuan",
			LastName:  "anh",
			Email:     "tuananh@example.com",
			Password:  "aaa123",
		},
	}
	for _, v := range users {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}
