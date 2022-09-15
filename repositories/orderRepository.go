package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrder() []domain.Order
	GetOrderByID(id int) (domain.Order, *helpers.AppError)
	CreateOrder(order domain.Order) domain.Order
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (repository *orderRepository) GetAllOrder() []domain.Order {
	var orders []domain.Order
	repository.db.Find(&orders)
	return orders
}

func (repository *orderRepository) GetOrderByID(id int) (domain.Order, *helpers.AppError) {
	var order domain.Order
	repository.db.First(&order, id)
	if order.OrderID == 0 {
		return order, helpers.NewNotFoundError("Order Not Found")
	} else {
		return order, nil
	}
}

func (repository *orderRepository) CreateOrder(order domain.Order) domain.Order {
	repository.db.Create(&order)
	return order
}
