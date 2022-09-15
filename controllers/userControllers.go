package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
}

type userController struct {
	userServices services.UserServices
	authServices services.AuthService
}

func NewUserController(userServices services.UserServices, authService services.AuthService) *userController {
	return &userController{userServices, authService}
}

func (controllers *userController) RegisterUser(c *gin.Context) {
	var regist dto.RegisterRequest
	err := c.ShouldBindJSON(&regist)
	if err != nil {
		response := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, errRegis := controllers.userServices.RegisterUser(regist)
	if errRegis != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helpers.APIResponse(http.StatusCreated, "Success", "Register User success!", newUser)

	c.JSON(http.StatusCreated, response)
}
