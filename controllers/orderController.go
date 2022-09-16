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
		response := helpers.APIResponse("Get Order Detail", http.StatusBadRequest, "Order Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := helpers.APIResponse("Get Order Detail", http.StatusOK, "Success", order)
		c.JSON(http.StatusOK, response)
	}
}

func (controller *orderController) CreateOrder(c *gin.Context) {
	order := dto.OrderRequest{}
	c.BindJSON(&order)
	controller.orderService.CreateOrder(order)
	response := helpers.APIResponse("Order Created", http.StatusCreated, "Success", order)
	c.JSON(http.StatusCreated, response)
}
