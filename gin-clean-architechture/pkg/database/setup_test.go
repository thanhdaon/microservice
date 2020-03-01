package database

import (
	"domain-driven-design/domain/entity"

	"github.com/jinzhu/gorm"
)

func NewTestDBConnection() *gorm.DB {
	DB_DRIVER := "postgres"
	TEST_DB_CONNECTION_STRING := "host=localhost port=5432 user=demo dbname=demo_test password=password sslmode=disable"
	db := NewDBConnection(DB_DRIVER, TEST_DB_CONNECTION_STRING)
	db.SingularTable(true)
	db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})
	return db
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
