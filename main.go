package main

import (
	"go-pos-api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.NewConnection()
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run("localhost:8080")
}
