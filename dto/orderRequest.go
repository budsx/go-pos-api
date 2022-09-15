package dto

import "time"

type OrderRequest struct {
	UserID       int       `json:"user_id"`
	CustomerName string    `json:"customer_name"`
	Amount       int       `json:"amount"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
