package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type DetailOrderService interface {
	CreateDetailOrder(request dto.OrderRequest, orderID int) *helpers.AppError
}

type detailOrderService struct {
	repository        repositories.DetailOrderRepository
	productRepository repositories.ProductRepository
}

func (service *detailOrderService) CreateDetailOrder(request dto.OrderRequest, orderID int) *helpers.AppError {
	detailOrder := []domain.DetailOrder{}

	for i, product := range request.Products {
		productDetail, _ := service.productRepository.GetProductById(product.ProductID)
		detailOrder[i].ProductID = product.ProductID
		detailOrder[i].OrderID = orderID
		detailOrder[i].Quantity = product.Quantity
		detailOrder[i].SubTotal = product.Quantity * productDetail.Price

		productDetail.Stock -= product.Quantity
		service.productRepository.UpdateProductById(productDetail, product.ProductID)
		service.repository.CreateDetailOrder(detailOrder[i])
	}

	return nil
}
