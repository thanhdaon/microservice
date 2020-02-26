package routes

import (
	"domain-driven-design/pkg/routes/api"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.POST("/signin", api.Signin)
	r.POST("/signup", api.Signup)

	return r
}
