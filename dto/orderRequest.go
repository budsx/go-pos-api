package dto

import "time"

type OrderRequest struct {
	UserID       int       `json:"user_id" binding:"required"`
	CustomerName string    `json:"customer_name" binding:"required"`
	Amount       int       `json:"amount" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
