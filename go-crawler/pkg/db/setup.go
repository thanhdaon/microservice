package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open("mysql", "admin:password@/email-db")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	log.Println("[INFO] mysql connected")
}

func Close() {
	db.Close()
}
