package controllers

import (
	"fmt"
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
	DeleteProductById(c *gin.Context)
	UpdateProductById(c *gin.Context)
	UploadImageProduct(c *gin.Context)
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

func (controller *productController) DeleteProductById(c *gin.Context) {
	productIdString := c.Param("product_id")
	productId, _ := strconv.Atoi(productIdString)
	_, err := controller.productService.DeleteProductById(productId)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Success delete Product", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (controller *productController) UpdateProductById(c *gin.Context) {
	productIdString := c.Param("product_id")
	productId, _ := strconv.Atoi(productIdString)
	var input dto.ProductRequest
	errShouldBindJSON := c.ShouldBindJSON(&input)
	if errShouldBindJSON != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	updatedProduct, err := controller.productService.UpdateProductById(input, productId)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Success Updated product", http.StatusOK, "success", updatedProduct)
	c.JSON(http.StatusOK, response)
}

func (controller *productController) UploadImageProduct(c *gin.Context) {
	file, err := c.FormFile("image")
	productIdString := c.Param("product_id")
	productId, _ := strconv.Atoi(productIdString)
	if err != nil {
		response := helpers.APIResponse("Error upload file", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	path := fmt.Sprintf("images/%s", file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helpers.APIResponse("Error upload file", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	product, _ := controller.productService.UploadImageProduct(productId, path)
	response := helpers.APIResponse("Success Upload image", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}
