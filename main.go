package main

import (
	"go-pos-api/config"
	"go-pos-api/controllers"
	"go-pos-api/middlewares"
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
	authMiddleware := middlewares.AuthMiddleware(userServices, authService)

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Connect Server with Gin Gonic",
		})
	})

	router.POST("/users", authMiddleware, userController.RegisterUser)
	router.POST("/login", userController.LoginUser)
	router.GET("/users", authMiddleware, userController.GetAllUsers)
	router.GET("/users/:user_id", authMiddleware, userController.GetUsersByID)
	router.PUT("/users/:user_id", authMiddleware, userController.UpdateUser)
	router.DELETE("/users/:user_id", authMiddleware, userController.DeleteUser)

	router.Run("localhost:8080")
}
