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
	GetAllUsers(c *gin.Context)
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
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", "Failed user already exist")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, errToken := controllers.authServices.GenerateToken(newUser.ID)
	if errToken != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", "Failed generate token")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Register user failed", "Failed create user")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userRegist := dto.RegisterInput(newUser, token)
	response := helpers.APIResponse(http.StatusCreated, "Success", "Register User success!", userRegist)
	c.JSON(http.StatusCreated, response)
}

func (controllers *userController) LoginUser(c *gin.Context) {
	request := dto.LoginRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse(http.StatusUnprocessableEntity, "Error", "Login user failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loginUser, errLogin := controllers.userServices.LoginUser(request)
	if errLogin != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Login user failed", "Wrong Password")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, errToken := controllers.authServices.GenerateToken(loginUser.ID)
	if errToken != nil {
		res := helpers.APIResponse(http.StatusBadRequest, "Error", "Login user failed", "Failed Generated Token")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userRequest := dto.UserInput(loginUser, token)
	response := helpers.APIResponse(http.StatusOK, "Success", "Login user success!", userRequest)
	c.JSON(http.StatusOK, response)
}

func (controllers *userController) GetAllUsers(c *gin.Context) {
	users, err := controllers.userServices.GetAllUsers()
	if err != nil {
		response := helpers.APIResponse(http.StatusInternalServerError, "Error", "GetAllUsers failed", "Failed GetAllUsers")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse(http.StatusOK, "Success", "GetAllUsers success", users)
	c.JSON(http.StatusOK, response)
}