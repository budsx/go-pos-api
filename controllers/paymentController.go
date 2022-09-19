package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	CreatePayment(c *gin.Context)
}

type paymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) *paymentController {
	return &paymentController{paymentService: paymentService}
}

func (controller *paymentController) CreatePayment(c *gin.Context) {
	var input dto.CreatePaymentInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	newPayment, err := controller.paymentService.CreatePayment(input)
	if err != nil {
		response := helpers.APIResponse("Something went wrong", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("Success Create Payment", http.StatusCreated, "success", newPayment)
	c.JSON(http.StatusCreated, response)
}
