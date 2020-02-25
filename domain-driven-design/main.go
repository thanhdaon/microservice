package main

import (
	"domain-driven-design/database"
	"domain-driven-design/domain/entity"
	"domain-driven-design/domain/usecase"
	"domain-driven-design/pkg/auth"

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

		authHelper = auth.NewAuthHelper()

		userUC = usecase.NewUserUsecase(userRepo, authHelper)
	)

	db.SingularTable(true)
	db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})

	token, err := userUC.Login("admin@gmail.com", "aaa123")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(token)

	defer db.Close()
}
