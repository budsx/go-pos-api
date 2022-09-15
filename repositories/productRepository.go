package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.Product, *helpers.AppError)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (repository *productRepository) GetAllProduct() ([]domain.Product, *helpers.AppError) {
	var err error
	var products []domain.Product
	if err = repository.db.Find(&products).Error; err != nil {
		helpers.Error("error db" + err.Error())
		return products, helpers.NewUnexpectedError("unExpected error" + err.Error())
	}
	return products, nil
}
