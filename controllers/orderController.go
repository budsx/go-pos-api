package controllers

import (
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetAllOrder(c *gin.Context)
}

type orderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{orderService}
}

func (controller *orderController) GetAllOrder(c *gin.Context) {
	orders := controller.orderService.GetAllOrder()
	response := helpers.APIResponse("Get All Orders", http.StatusOK, "Success", orders)
	c.JSON(http.StatusOK, response)
}
