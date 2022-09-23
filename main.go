package main

import (
	"go-pos-api/config"
	"go-pos-api/controllers"
	"go-pos-api/middlewares"
	"go-pos-api/repositories"
	"go-pos-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewConnection()
	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)
	detailRepository := repositories.NewDetailOrderRepository(db)

	orderService := services.NewOrderService(orderRepository, detailRepository, productRepository)
	orderController := controllers.NewOrderController(orderService)
	userServices := services.NewUserService(userRepository)
	authService := services.NewService()
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)
	userController := controllers.NewUserController(userServices, authService)
	authMiddleware := middlewares.AuthMiddleware(userServices, authService)
	productMiddleware := middlewares.ProductMiddleware(userServices, authService)
	orderMiddleware := middlewares.OrderMiddleware(userServices, authService)

	router := gin.Default()

	router.POST("/users", authMiddleware, userController.RegisterUser)
	router.POST("/login", userController.LoginUser)
	router.GET("/users", authMiddleware, userController.GetAllUsers)
	router.GET("/users/:user_id", authMiddleware, userController.GetUsersByID)
	router.PUT("/users/:user_id", authMiddleware, userController.UpdateUser)
	router.DELETE("/users/:user_id", authMiddleware, userController.DeleteUser)

	router.GET("/products", productController.GetAllProduct)
	router.GET("/products/:product_id", productMiddleware, productController.GetProductById)
	router.DELETE("/products/:product_id", productMiddleware, productController.DeleteProductById)
	router.PUT("/products/:product_id", productMiddleware, productController.UpdateProductById)
	router.POST("/products", productMiddleware, productController.CreateProduct)
	router.POST("/products-image/:product_id", productController.UploadImageProduct)

	router.GET("/orders", orderMiddleware, orderController.GetAllOrder)
	router.GET("/orders/:order_id", orderMiddleware, orderController.GetOrderByID)
	router.POST("/orders", orderMiddleware, orderController.CreateOrder)

	router.Run("localhost:8080")
}
