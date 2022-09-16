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
	LoginUser(c *gin.Context)
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
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, errRegis := controllers.userServices.RegisterUser(regist)
	if errRegis != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", "Failed add user")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helpers.APIResponse(http.StatusCreated, "Success", "Register User success!", newUser)

	c.JSON(http.StatusCreated, response)
}

func (controllers *userController) LoginUser(c *gin.Context) {
	request := dto.LoginRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "error", "Login user failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loginUser, _ := controllers.userServices.LoginUser(request)

	token, errToken := controllers.authServices.GenerateToken(loginUser.ID)
	if errToken != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "error", "Login user failed", nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userRequest := dto.UserInput(loginUser, token)
	response := helpers.APIResponse(http.StatusOK, "success", "Login user success!", userRequest)
	c.JSON(http.StatusOK, response)
}
