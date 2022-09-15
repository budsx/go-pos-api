package dto

import "time"

type OrderResponse struct {
	OrderID      int       `json:"order_id"`
	UserID       int       `json:"user_id"`
	CustomerName string    `json:"customer_name"`
	Amount       int       `json:"amount"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
