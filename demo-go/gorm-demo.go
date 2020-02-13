package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Person struct {
	ID   int
	Name string
	Age  int
	City string
}

func main() {
	setupDB()
	initTables()
	CRUD()
	defer db.Close()
}

func setupDB() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=demo dbname=demo password=demo_password sslmode=disable")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db.SingularTable(true)
}

func initTables() {
	db.DropTableIfExists(&Person{})
	db.AutoMigrate(&Person{})
}

func CRUD() {
	// createDemo()
	queryDemo()
}

func createDemo() {
	person := Person{Name: "thanh", Age: 21, City: "HaNoi"}
	fmt.Println(db.NewRecord(person))
	db.Create(&person)
	fmt.Println(db.NewRecord(person))
}

func queryDemo() {
	john := Person{Name: "john", Age: 21, City: "HaNoi"}
	maven := Person{Name: "maven", Age: 21, City: "HaNoi"}
	otis := Person{Name: "otis", Age: 21, City: "HaNoi"}
	ruby := Person{Name: "ruby", Age: 21, City: "HaNoi"}

	db.Create(&john)
	db.Create(&maven)
	db.Create(&otis)
	db.Create(&ruby)

	var persons []Person
	db.Where("city = ?", "HaNoi").Find(&persons)
	fmt.Println(persons)

	var person Person
	db.First(&person, 2)
	fmt.Println(person)
}
