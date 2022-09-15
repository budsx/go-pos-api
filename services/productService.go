package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type ProductService interface {
	GetAllProduct() ([]dto.ProductResponse, *helpers.AppError)
	GetProductById(int) (dto.ProductResponse, *helpers.AppError)
	CreateProduct(dto.ProductRequest) (dto.ProductResponse, *helpers.AppError)
	DeleteProductById(int) (domain.Product, *helpers.AppError)
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

func (service *productService) GetProductById(id int) (dto.ProductResponse, *helpers.AppError) {
	product, err := service.productService.GetProductById(id)
	productResponse := dto.ProductResponse{}
	if err != nil {
		helpers.Error("error" + err.Message)
		return productResponse, err
	} else {
		productResponse.ID = product.ID
		productResponse.MerchantId = product.MerchantId
		productResponse.Name = product.Name
		productResponse.Price = product.Price
		productResponse.Stock = product.Stock
		return productResponse, nil
	}
}

func (service *productService) CreateProduct(input dto.ProductRequest) (dto.ProductResponse, *helpers.AppError) {
	p := domain.Product{}
	p.Name = input.Name
	p.Price = input.Price
	p.Stock = input.Stock
	p.MerchantId = input.MerchantId
	product, err := service.productService.CreateProduct(p)
	productResponse := dto.ProductResponse{}
	if err != nil {
		helpers.Error("error" + err.Message)
		return productResponse, err
	} else {
		productResponse.ID = product.ID
		productResponse.MerchantId = product.MerchantId
		productResponse.Name = product.Name
		productResponse.Price = product.Price
		productResponse.Stock = product.Stock
		return productResponse, nil
	}
}

func (service *productService) DeleteProductById(id int) (domain.Product, *helpers.AppError) {
	product, err := service.productService.DeleteProductById(id)
	if err != nil {
		helpers.Error("Error" + err.Message)
		return product, err
	}
	return product, nil
}
