package main

import (
	"go-pos-api/config"
	"go-pos-api/controllers"
	"go-pos-api/repositories"
	"go-pos-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewConnection()
	router := gin.Default()

	// repository
	productRepository := repositories.NewProductRepository(db)
	paymentRepository := repositories.NewPaymentRepository(db)

	// service
	productService := services.NewProductService(productRepository)
	midtransService := services.NewMidTransService(paymentRepository)
	paymentService := services.NewPaymentService(paymentRepository, midtransService)

	// controller
	productController := controllers.NewProductController(productService)
	paymentController := controllers.NewPaymentController(paymentService, midtransService)

	router.GET("/products", productController.GetAllProduct)
	router.GET("/products/:product_id", productController.GetProductById)
	router.DELETE("/products/:product_id", productController.DeleteProductById)
	router.PUT("/products/:product_id", productController.UpdateProductById)
	router.POST("/products", productController.CreateProduct)
	// router.POST("/products-image/:product_id", productController.UploadImageProduct)
	router.POST("/payments", paymentController.CreatePayment)
	router.POST("/payments/notification", paymentController.GetNotificationFromMidtrans)

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Connect Server with Gin Gonic",
		})
	})

	router.Run("localhost:8080")
}
