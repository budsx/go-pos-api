package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionController struct {
	userServices services.UserServices
}

func NewSessionController(userServices services.UserServices) *sessionController {
	return &sessionController{userServices}
}

func (controllers *sessionController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "sesion.html", nil)
}

func (controllers *sessionController) Create(c *gin.Context) {
	var input dto.LoginRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helpers.APIResponse("Login user failed", http.StatusBadRequest, "Error", "Wrong Password")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, _ := controllers.userServices.LoginUser(input)
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (controllers *sessionController) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
