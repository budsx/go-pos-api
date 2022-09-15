package services

import (
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type ProductService interface {
	GetAllProduct() ([]dto.ProductResponse, *helpers.AppError)
}

type productService struct {
	productService repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (service *productService) GetAllProduct() ([]dto.ProductResponse, *helpers.AppError) {
	products, err := service.productService.GetAllProduct()
	productResponse := []dto.ProductResponse{}
	if err != nil {
		return productResponse, err
	}
	for _, product := range products {
		productResponse = append(productResponse, dto.ProductResponse{
			ID:         product.ID,
			MerchantId: product.MerchantId,
			Price:      product.Price,
			Stock:      product.Stock,
			Name:       product.Name,
		})
	}
	return productResponse, nil
}
