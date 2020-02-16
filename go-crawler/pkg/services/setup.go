package services

import (
	"email-crawler/pkg/database"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {
	db = database.DB
}
