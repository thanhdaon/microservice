package main

import (
	"domain-driven-design/database"
	"domain-driven-design/domain/usecase"
	"domain-driven-design/pkg/auth"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		DB_DRIVER            = os.Getenv("DB_DRIVER")
		DB_CONNECTION_STRING = os.Getenv("DB_CONNECTION_STRING")

		db       = database.NewDBConnection(DB_DRIVER, DB_CONNECTION_STRING)
		userRepo = database.NewUserRepository(db)

		authUtil = auth.NewAuthUtil()

		userUC = usecase.NewUserUsecase(userRepo)
	)

	fmt.Println(authUtil, userUC)

	defer db.Close()
}
