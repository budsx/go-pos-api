package main

import (
	"go-pos-api/config"
	"go-pos-api/controllers"
	"go-pos-api/repositories"
	"go-pos-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := config.NewConnection()
	godotenv.Load()
	userRepository := repositories.NewUserRepository(db)
	userServices := services.NewUserService(userRepository)
	authService := services.NewService()
	userController := controllers.NewUserController(userServices, authService)

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Connect Server with Gin Gonic",
		})
	})

	router.POST("/login", userController.Login)
	router.POST("/register", userController.RegisterUser)
	router.Run("localhost:8080")
}
