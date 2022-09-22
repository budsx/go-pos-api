package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetUsersByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
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
		response := helpers.APIResponse("Register user failed", http.StatusBadRequest, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, errRegis := controllers.userServices.RegisterUser(regist)
	if errRegis != nil {
		res := helpers.APIResponse("Register user failed", http.StatusBadRequest, "Error", "Failed user already exist")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, errToken := controllers.authServices.GenerateToken(newUser.ID)
	if errToken != nil {
		res := helpers.APIResponse("Register user failed", http.StatusBadRequest, "Error", "Failed generate token")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userRegist := dto.RegisterInput(newUser, token)
	response := helpers.APIResponse("Register User success!", http.StatusCreated, "Success", userRegist)
	c.JSON(http.StatusCreated, response)
}

func (controllers *userController) LoginUser(c *gin.Context) {
	request := dto.LoginRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("Login user failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loginUser, errLogin := controllers.userServices.LoginUser(request)
	if errLogin != nil {
		res := helpers.APIResponse("Login user failed", http.StatusBadRequest, "Error", "Wrong Password")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, errToken := controllers.authServices.GenerateToken(loginUser.ID)
	if errToken != nil {
		res := helpers.APIResponse("Login user failed", http.StatusBadRequest, "Error", "Failed Generated Token")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	userRequest := dto.UserInput(loginUser, token)
	response := helpers.APIResponse("Login user success!", http.StatusOK, "Success", userRequest)
	c.JSON(http.StatusOK, response)
}

func (controllers *userController) GetAllUsers(c *gin.Context) {
	users, err := controllers.userServices.GetAllUsers()
	if err != nil {
		response := helpers.APIResponse("GetAllUsers failed", http.StatusInternalServerError, "Error", "Failed GetAllUsers")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("GetAllUsers success", http.StatusOK, "Success", users)
	c.JSON(http.StatusOK, response)
}

func (controllers *userController) GetUsersByID(c *gin.Context) {
	usersIdString := c.Param("user_id")
	usersID, _ := strconv.Atoi(usersIdString)
	users, err := controllers.userServices.GetUsersByID(usersID)
	if err != nil {
		response := helpers.APIResponse("GetUsersByID failed", http.StatusInternalServerError, "Error", "Failed User ID not found")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("GetUsersByID success", http.StatusOK, "Success", users)
	c.JSON(http.StatusOK, response)
}

func (controllers *userController) UpdateUser(c *gin.Context) {
	usersIdString := c.Param("user_id")
	usersID, _ := strconv.Atoi(usersIdString)
	var request dto.RegisterRequest
	errShouldBindJSON := c.ShouldBindJSON(&request)

	if errShouldBindJSON != nil {
		response := helpers.APIResponse("Update User failed", http.StatusInternalServerError, "Error", "Failed Update User")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	updatedUsers, err := controllers.userServices.UpdateUser(request, usersID)
	if err != nil {
		response := helpers.APIResponse("Update User failed", http.StatusInternalServerError, "Error", "Failed Update User")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Update User success", http.StatusOK, "Success", updatedUsers)
	c.JSON(http.StatusOK, response)
}

func (controllers *userController) DeleteUser(c *gin.Context) {
	usersIdString := c.Param("user_id")
	usersID, _ := strconv.Atoi(usersIdString)
	_, err := controllers.userServices.DeleteUser(usersID)
	if err != nil {
		response := helpers.APIResponse("Delete User failed", http.StatusInternalServerError, "Error", "Failed Delete User")
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Delete User success", http.StatusOK, "Success", "Success Delete User Where ID = "+usersIdString)
	c.JSON(http.StatusOK, response)
}
