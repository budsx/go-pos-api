package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type DetailOrderRepository interface {
	CreateDetailOrder(detail domain.DetailOrder) (domain.DetailOrder, *helpers.AppError)
	GetDetailOrder(int) ([]domain.DetailOrder, *helpers.AppError)
}

type detailOrderRepository struct {
	db *gorm.DB
}

func NewDetailOrderRepository(db *gorm.DB) DetailOrderRepository {
	return &detailOrderRepository{db}
}

func (repository *detailOrderRepository) CreateDetailOrder(detail domain.DetailOrder) (domain.DetailOrder, *helpers.AppError) {
	err := repository.db.Create(&detail)
	if err != nil {
		helpers.NewBadRequestError("Bad Request")
	}
	return detail, nil
}

func (repository *detailOrderRepository) GetDetailOrder(orderId int) ([]domain.DetailOrder, *helpers.AppError) {
	detaiOrder := []domain.DetailOrder{}
	err := repository.db.Where("order_id = ?", orderId).Find(&detaiOrder).Error
	if err != nil {
		return detaiOrder, helpers.NewUnexpectedError("unexpected error")
	}
	return detaiOrder, nil
}
