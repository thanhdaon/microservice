package crawler

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Domain struct {
	ID     int
	Domain string
}

type Email struct {
	ID            int
	Email         string
	ResourceCount int         `gorm:"default:1"`
	Resources     []*Resource `gorm:"many2many:email_resource"`
}

type Resource struct {
	ID       int
	Resource string
	Emails   []*Email `gorm:"many2many:email_resource"`
}

func SetupDB() {
	var err error
	db, err = gorm.Open("mysql", "demo:password@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	db.SingularTable(true)
	initTables()
	log.Println("[INFO] mysql setup")
}

func CleanupDB() {
	db.Close()
}

func initTables() {
	db.DropTableIfExists(&Domain{})
	db.DropTableIfExists(&Email{})
	db.DropTableIfExists(&Resource{})
	db.AutoMigrate(&Domain{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Resource{})
}
