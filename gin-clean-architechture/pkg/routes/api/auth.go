package api

import (
	"domain-driven-design/domain/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signinReqBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signin(c *gin.Context) {
	var json signinReqBody

	if err := c.BindJSON(&json); err != nil {
		return
	}

	token, err := authUC.Signin(json.Email, json.Password)
	if err != nil {
		if err == e.WRONG_PASSWORD || err == e.USER_NOT_FOUND {
			c.JSON(http.StatusOK, response{
				Ok:   false,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response{
			Ok:   false,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Ok:   true,
		Msg:  "",
		Data: token,
	})
	return
}

type signupReqBody struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var json signupReqBody

	if err := c.BindJSON(&json); err != nil {
		return
	}

	user, err := authUC.Signup(json.Email, json.Password, json.FirstName, json.LastName)
	if err != nil {
		c.JSON(http.StatusOK, response{
			Ok:   false,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Ok:   true,
		Msg:  "",
		Data: user.ID,
	})
	return
}
