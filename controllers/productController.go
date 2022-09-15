package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProduct(c *gin.Context)
	GetProductById(c *gin.Context)
	CreateProduct(c *gin.Context)
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
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
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
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Success Get Product By Id"+productIdString, http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (controller *productController) CreateProduct(c *gin.Context) {
	var input dto.ProductRequest
	errShouldBindJSON := c.ShouldBindJSON(&input)
	if errShouldBindJSON != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	product, err := controller.productService.CreateProduct(input)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Create Product Success", http.StatusCreated, "success", product)
	c.JSON(http.StatusCreated, response)
}
