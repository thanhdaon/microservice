package routes

import (
	"domain-driven-design/pkg/routes/api"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.POST("/login", api.Login)

	return r
}
