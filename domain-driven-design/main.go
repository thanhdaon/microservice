package main

import (
	"domain-driven-design/database"
	"domain-driven-design/domain/usecase"
	"log"
	"os"
)

var (
	DB_DRIVER            = os.Getenv("DB_DRIVER")
	DB_CONNECTION_STRING = os.Getenv("DB_CONNECTION_STRING")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var (
		db       = database.NewDBConnection(DB_DRIVER, DB_CONNECTION_STRING)
		userRepo = database.NewUserRepository(db)
		userUC   = usecase.NewUserUsecase(userRepo)
	)
	defer db.Close()
}
