package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBConnection(dbDriver, dbConnectionString string) *gorm.DB {
	db, err := gorm.Open(dbDriver, dbConnectionString)
	if err != nil {
		log.Fatalln("Could not connect to DB! \n", err)
	}
	log.Println("DB connected")
	return db
}
