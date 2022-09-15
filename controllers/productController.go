package controllers

import (
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProduct(c *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *productController {
	return &productController{productService: productService}
}

func (controller *productController) GetAllProduct(c *gin.Context) {
	products, err := controller.productService.GetAllProduct()
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}
	response := helpers.APIResponse("Success Get All Product", http.StatusOK, "error", products)
	c.JSON(http.StatusOK, response)
}

func (controller *productController) GetProductById(c *gin.Context) {
	productIdString := c.Param("product_id")
	productId, _ := strconv.Atoi(productIdString)
	product, err := controller.productService.GetProductById(productId)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadGateway, response)
		return
	}
	response := helpers.APIResponse("Success Get Product By Id"+productIdString, http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}
