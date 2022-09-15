package repositories

import (
	"go-pos-api/domain"
	"go-pos-api/helpers"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]domain.Product, *helpers.AppError)
	GetProductById(int) (domain.Product, *helpers.AppError)
	CreateProduct(domain.Product) (domain.Product, *helpers.AppError)
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

func (repository *productRepository) GetProductById(id int) (domain.Product, *helpers.AppError) {
	var err error
	var product domain.Product
	if err = repository.db.Where("product_id = ?", id).Find(&product).Error; err != nil {
		helpers.Error("error db" + err.Error())
		return product, helpers.NewUnexpectedError("unExpected error" + err.Error())
	}
	return product, nil
}

func (repository *productRepository) CreateProduct(product domain.Product) (domain.Product, *helpers.AppError) {
	var err error
	if err = repository.db.Create(&product).Error; err != nil {
		helpers.Error("unexpected error" + err.Error())
		return product, helpers.NewUnexpectedError("unexpected error")
	}
	return product, nil
}
