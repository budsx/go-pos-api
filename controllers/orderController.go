package controllers

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetAllOrder(c *gin.Context)
	GetOrderByID(c *gin.Context)
	CreateOrder(c *gin.Context)
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

func (controller *orderController) GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("order_id"))
	order, err := controller.orderService.GetOrderByID(id)
	if err != nil {
		response := helpers.APIResponse("Order Not Found", http.StatusBadRequest, "Order Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.APIResponse("Get Order Detail", http.StatusOK, "Success", order)
		c.JSON(http.StatusOK, response)
		return
	}
}

func (controller *orderController) CreateOrder(c *gin.Context) {
	request := dto.OrderRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response := helpers.APIResponse("Failed to create new order", http.StatusBadRequest, "Error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	order, errs := controller.orderService.CreateOrder(request)
	if errs != nil {
		response := helpers.APIResponse("Failed to create new order", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.APIResponse("Order Created", http.StatusCreated, "Success", order)
		c.JSON(http.StatusCreated, response)
		return
	}

}
