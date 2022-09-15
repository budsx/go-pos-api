package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userServices services.UserServices
	authServices services.AuthService
}

func NewUserController(userServices services.UserServices, authService services.AuthService) *userController {
	return &userController{userServices, authService}
}

func (controllers *userController) RegisterUser(c *gin.Context) {
	var input dto.UserDTO
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Register user failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := controllers.userServices.RegisterUser(input)
	if err != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "error", "Register user failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, err := controllers.authServices.GenerateToken(newUser.ID)
	if err != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "error", "Register account failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "error", "Register user failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	createUser := dto.UserInput(newUser, token)
	response := helpers.APIResponse(http.StatusCreated, "success", "Register member success!", createUser)

	c.JSON(http.StatusCreated, response)
}

func (controllers *userController) Login(c *gin.Context) {
	var input dto.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Login member failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loginUser, err := controllers.userServices.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Login member failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := controllers.authServices.GenerateToken(loginUser.ID)
	if err != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "error", "Login member failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userInput := dto.UserInput(loginUser, token)
	response := helpers.APIResponse(http.StatusOK, "success", "Login member success!", userInput)
	c.JSON(http.StatusOK, response)
}
