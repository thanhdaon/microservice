package app

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Server *echo.Echo
	API    *echo.Group
)

func init() {
	Server = echo.New()

	API = Server.Group("/api")

	Server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	Server.Use(middleware.Recover())
}

func Start(port int) {
	Server.Logger.Fatal(Server.Start(fmt.Sprintf(":%d", port)))
}
