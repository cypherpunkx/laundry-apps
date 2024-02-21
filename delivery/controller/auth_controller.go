package controller

import (
	"net/http"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router *gin.Engine
	authUc usecase.AuthUseCase
}

func (a *AuthController) createHandler(c *gin.Context) {
	var payload model.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error" : err.Error(),
		})
		return
	}
	token, err := a.authUc.Login(payload.Username, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H {
		"message" : "Success Login",
		"token" : token,
	})
}


func NewAuthController(r *gin.Engine,authUseCase usecase.AuthUseCase) {
	controller := AuthController {
		router: r,
		authUc: authUseCase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/login",controller.createHandler)
}