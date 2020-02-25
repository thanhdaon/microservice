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
		JWT_SECRET           = os.Getenv("JWT_SECRET")

		db       = database.NewDBConnection(DB_DRIVER, DB_CONNECTION_STRING)
		userRepo = database.NewUserRepository(db)

		authHelper = auth.NewAuthHelper([]byte(JWT_SECRET))

		userUC = usecase.NewUserUsecase(userRepo, authHelper)
	)

	db.SingularTable(true)
	// db.DropTableIfExists(&entity.User{})
	db.AutoMigrate(&entity.User{})

	user, err := userUC.Signup("adminm@gmail.com", "aaa123", "Thanh", "dao")
	if err != nil {
		log.Println(err)
	}
	log.Println(user)

	token, err := userUC.Signin("admin@gmail.com", "aaa123")
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(token)

	defer db.Close()
}
