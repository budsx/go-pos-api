package repositories

import (
	"go-pos-api/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrder() []domain.Order
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
