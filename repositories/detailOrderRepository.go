package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type DetailOrderRepository interface {
	CreateDetailOrder(detail domain.DetailOrder) (domain.DetailOrder, *helpers.AppError)
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
