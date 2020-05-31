package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_URI = "4wA9RL0rTf:LoTYJa36UN@(remotemysql.com)/4wA9RL0rTf?charset=utf8&parseTime=True&loc=Local"

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`
}

type Product struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"-"`
	CategoryID  int    `json:"category_id"`
	Quantity    int    `json:"quantity"`
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Price       string `gorm:"type:varchar(255);unique;not null" json:"price"`
	ImageUrl    string `gorm:"type:varchar(255);unique;not null" json:"image"`
	Description string `gorm:"type:varchar(255);unique;not null" json:"description"`
}

func main() {
	db, err := gorm.Open("mysql", DB_URI)
	check(err)
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})

	seed(db)
}

func seed(db *gorm.DB) {
	jsonFile, err := os.Open("static/adapter-cu-sac.json")
	check(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	check(err)

	var products []*Product
	err = json.Unmarshal(byteValue, &products)
	check(err)

	for _, p := range products {
		fmt.Println("--------------")
		fmt.Println(p.Name)
		fmt.Println(p.ImageUrl)
		fmt.Println(p.Price)
		fmt.Println(p.ID)
		db.Save(p)
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
