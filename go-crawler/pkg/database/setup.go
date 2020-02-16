package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Setup() {
	var err error
	DB, err = gorm.Open("mysql", "admin:password@/email-db")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	DB.SingularTable(true)
	initTables()
	log.Println("[INFO] mysql setup")
}

func initTables() {
	DB.DropTableIfExists(&Domain{})
	DB.DropTableIfExists(&Email{})
	DB.DropTableIfExists(&Resource{})
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&Email{})
	DB.AutoMigrate(&Resource{})
}

func Close() {
	DB.Close()
}
