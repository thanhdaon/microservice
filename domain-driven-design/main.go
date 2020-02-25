package main

import (
	"domain-driven-design/domain/usecase"
	"domain-driven-design/pkg/auth"
	"domain-driven-design/pkg/database"
	"domain-driven-design/pkg/routes/api"

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
	defer db.Close()

	api.Setup(api.Dependences{UserUC: userUC})

	TestUserUC(userUC)
}

func TestUserUC(userUC usecase.UserUsecase) {
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
}
