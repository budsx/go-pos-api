package dto

type CreatePaymentInput struct {
	Amount          int    `json:"amount" binding:"required"`
	OrderID         int    `json:"order_id" binding:"required"`
	PaymentCategory string `json:"payment_category" binding:"required"`
	Status          string `json:"payment_status" binding:"required"`
}
