package services

import (
	"go-pos-api/domain"
	"go-pos-api/dto"
	"go-pos-api/helpers"
	"go-pos-api/repositories"
)

type OrderService interface {
	GetAllOrder() []dto.ListOrderAndOrderDetailResponse
	GetOrderByID(id int) (dto.ListOrderAndOrderDetailResponse, *helpers.AppError)
	CreateOrder(request dto.OrderRequest) (dto.OrderResponse, *helpers.AppError)
}

type orderService struct {
	orderRepository       repositories.OrderRepository
	detailOrderRepository repositories.DetailOrderRepository
	productRepository     repositories.ProductRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, detailOrderRepository repositories.DetailOrderRepository, productRepository repositories.ProductRepository) OrderService {
	return &orderService{orderRepository, detailOrderRepository, productRepository}
}

func (service *orderService) GetAllOrder() []dto.ListOrderAndOrderDetailResponse {
	orders := service.orderRepository.GetAllOrder()

	var orderResponses []dto.ListOrderAndOrderDetailResponse
	for _, order := range orders {
		productsOrder, _ := service.detailOrderRepository.GetDetailOrder(order.OrderID)
		productsOrders := []dto.ProductsOrder{}
		for _, product := range productsOrder {
			getProduct, _ := service.productRepository.GetProductById(product.ProductID)
			productsOrders = append(productsOrders, dto.ProductsOrder{
				ProductID:   product.ProductID,
				Quantity:    product.Quantity,
				ProductName: getProduct.Name,
				Subtotal:    getProduct.Price * product.Quantity,
				Price:       getProduct.Price,
			})
		}
		orderResponses = append(orderResponses, dto.ListOrderAndOrderDetailResponse{
			OrderID:      order.OrderID,
			UserID:       order.UserID,
			CustomerName: order.CustomerName,
			Amount:       order.Amount,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			Products:     productsOrders,
		})
	}
	return orderResponses
}

func (service *orderService) GetOrderByID(id int) (dto.ListOrderAndOrderDetailResponse, *helpers.AppError) {
	order, err := service.orderRepository.GetOrderByID(id)
	if err != nil {
		return dto.ListOrderAndOrderDetailResponse{}, helpers.NewUnexpectedError("Internal Server Error")
	}
	productsOrder, err := service.detailOrderRepository.GetDetailOrder(id)
	orderProducts := []dto.ProductsOrder{}
	for _, product := range productsOrder {
		getProduct, _ := service.productRepository.GetProductById(product.ProductID)
		orderProducts = append(orderProducts, dto.ProductsOrder{
			ProductID:   product.ProductID,
			ProductName: getProduct.Name,
			Quantity:    product.Quantity,
			Price:       getProduct.Price,
			Subtotal:    getProduct.Price * product.Quantity,
		})
	}
	if err != nil {
		return dto.ListOrderAndOrderDetailResponse{}, helpers.NewUnexpectedError("Internal Server Error")
	}
	return dto.ListOrderAndOrderDetailResponse{
		OrderID:      order.OrderID,
		UserID:       order.UserID,
		CustomerName: order.CustomerName,
		Amount:       order.Amount,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
		Products:     orderProducts,
	}, nil
}

func (service *orderService) CreateOrder(request dto.OrderRequest) (dto.OrderResponse, *helpers.AppError) {
	var order domain.Order

	order.UserID = request.UserID
	order.CustomerName = request.CustomerName
	var grandTotal int

	order, err := service.orderRepository.CreateOrder(order)
	if err != nil {
		helpers.NewBadRequestError("Bad Request")
	}

	detailOrder := []domain.DetailOrder{}
	var orderProducts []dto.OrderProductRequest

	counter := 0

	for _, v := range request.Products {
		productDetail, _ := service.productRepository.GetProductById(v.ProductID)
		if productDetail.Stock-v.Quantity < 1 {
			counter -= 1
		} else {
			counter += 1
		}
	}

	if counter < 1 {
		return dto.OrderResponse{}, helpers.NewBadRequestError("Bad Request Error")
	} else {
		for i, product := range request.Products {
			productDetail, _ := service.productRepository.GetProductById(product.ProductID)
			detailOrder = append(detailOrder, domain.DetailOrder{
				ProductID: product.ProductID,
				OrderID:   order.OrderID,
				Quantity:  product.Quantity,
				SubTotal:  product.Quantity * productDetail.Price,
			})
			grandTotal += product.Quantity * productDetail.Price
			productDetail.Stock -= product.Quantity
			tempProduct, _ := service.productRepository.UpdateProductById(productDetail, product.ProductID)
			orderProducts = append(orderProducts, dto.OrderProductRequest{
				ProductID:   tempProduct.ID,
				ProductName: tempProduct.Name,
				Quantity:    product.Quantity,
			})
			service.detailOrderRepository.CreateDetailOrder(detailOrder[i])
		}
		return dto.OrderResponse{
			OrderID:      order.OrderID,
			UserID:       order.UserID,
			CustomerName: order.CustomerName,
			Amount:       grandTotal,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			Products:     orderProducts,
		}, nil
	}
}
