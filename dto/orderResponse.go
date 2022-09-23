package dto

import (
	"time"
)

type OrderResponse struct {
	OrderID      int                   `json:"order_id"`
	UserID       int                   `json:"user_id"`
	CustomerName string                `json:"customer_name"`
	Amount       int                   `json:"amount"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	Products     []OrderProductRequest `json:"products"`
}

type ListOrderAndOrderDetailResponse struct {
	OrderID      int             `json:"order_id"`
	UserID       int             `json:"user_id"`
	CustomerName string          `json:"customer_name"`
	Amount       int             `json:"amount"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Products     []ProductsOrder `json:"products"`
}

type ProductsOrder struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Subtotal    int    `json:"subtotal"`
}
