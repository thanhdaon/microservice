package database

import (
	"domain-driven-design/domain/entity"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBConnection(dbDriver, dbConnectionString string) *gorm.DB {
	db, err := gorm.Open(dbDriver, dbConnectionString)
	if err != nil {
		log.Fatalln("Could not connect to DB! \n", err)
	}
	autoMigrate(db)
	log.Println("DB connected")
	return db
}

func autoMigrate(db *gorm.DB) {
	db.SingularTable(true)
	db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})
}
