package http

import (
	"demo/pkg/app"
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	app.API.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!\n",
		})
	})
}
