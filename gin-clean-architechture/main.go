package main

import (
	"domain-driven-design/domain/usecase"
	"domain-driven-design/pkg/auth"
	"domain-driven-design/pkg/database"
	"domain-driven-design/pkg/routes"
	"domain-driven-design/pkg/routes/api"
	"net/http"
	"time"

	"log"
	"os"
)

var (
	DB_DRIVER            = os.Getenv("DB_DRIVER")
	DB_CONNECTION_STRING = os.Getenv("DB_CONNECTION_STRING")
	JWT_SECRET           = os.Getenv("JWT_SECRET")

	db       = database.NewDBConnection(DB_DRIVER, DB_CONNECTION_STRING)
	userRepo = database.NewUserRepository(db)

	authHelper = auth.NewAuthHelper([]byte(JWT_SECRET))

	authUC = usecase.NewUserUsecase(userRepo, authHelper)
)

func main() {
	defer db.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	api.Setup(api.Dependences{AuthUC: authUC})
	server := &http.Server{
		Addr:           ":8000",
		Handler:        routes.New(),
		ReadTimeout:    6 * time.Second,
		WriteTimeout:   6 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
