package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

func SetupDB(dbdriver, dbConnectionString string) *gorm.DB {
	db, err := gorm.Open(dbdriver, dbConnectionString)
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	log.Println("DB connected")
	return db
}
