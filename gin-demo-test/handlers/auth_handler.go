package handlers

import (
	"gin-demo-test/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, response{
		Ok:   true,
		Msg:  "",
		Data: &entity.User{},
	})
}
