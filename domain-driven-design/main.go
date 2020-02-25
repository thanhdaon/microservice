package main

import (
	"domain-driven-design/database"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbdriver := os.Getenv("DB_DRIVER")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db := database.NewDBConnection(dbdriver, dbConnectionString)
	defer db.Close()
}
