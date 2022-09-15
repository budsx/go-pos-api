package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
	"time"
)

type OrderService interface {
	GetAllOrder() []dto.OrderResponse
	GetOrderByID(id int) (dto.OrderResponse, *helpers.AppError)
	CreateOrder(request dto.OrderRequest) dto.OrderResponse
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &orderService{orderRepository}
}

func (service *orderService) GetAllOrder() []dto.OrderResponse {
	orders := service.orderRepository.GetAllOrder()

	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, dto.OrderResponse{
			OrderID:      order.OrderID,
			UserID:       order.UserID,
			CustomerName: order.CustomerName,
			Amount:       order.Amount,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
		})
	}
	return orderResponses
}

func (service *orderService) GetOrderByID(id int) (dto.OrderResponse, *helpers.AppError) {
	order, err := service.orderRepository.GetOrderByID(id)
	if err != nil {
		return dto.OrderResponse{}, helpers.NewUnexpectedError("Internal Server Error")
	}
	return dto.OrderResponse{
		OrderID:      order.OrderID,
		UserID:       order.UserID,
		CustomerName: order.CustomerName,
		Amount:       order.Amount,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}, nil
}

func (service *orderService) CreateOrder(request dto.OrderRequest) dto.OrderResponse {
	var order domain.Order

	order.UserID = request.UserID
	order.CustomerName = request.CustomerName
	order.Amount = request.Amount
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	order = service.orderRepository.CreateOrder(order)

	return dto.OrderResponse{
		OrderID:      order.OrderID,
		UserID:       order.UserID,
		CustomerName: order.CustomerName,
		Amount:       order.Amount,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}
