package repositories

import (
	"errors"
	"go-pos-api/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrder() []domain.Order
	GetOrderByID(id int) (domain.Order, error)
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

func (repository *orderRepository) GetOrderByID(id int) (domain.Order, error){
	var order domain.Order
	repository.db.First(&order, id)
	if order.OrderID == 0 {
		return order, errors.New("Order Not Found")
	} else {
		return order, nil
	}
}