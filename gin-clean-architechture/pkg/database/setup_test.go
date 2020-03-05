package database

import (
	"domain-driven-design/domain/entity"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewTestDBConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=demo dbname=demo_test password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&entity.User{})
	return db
}
