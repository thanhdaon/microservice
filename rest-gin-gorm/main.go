package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"rest-gin-gorm/product"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=demo dbname=demo password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.DropTableIfExists(&product.Product{})
	db.AutoMigrate(&product.Product{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
