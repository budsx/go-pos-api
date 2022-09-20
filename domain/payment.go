package domain

import "time"

type Payment struct {
	ID              int    `gorm:"column:payment_id"`
	OrderID         int    `gorm:"column:order_id"`
	Status          string `gorm:"column:status"`
	Amount          int    `gorm:"column:amount"`
	PaymentCategory string `gorm:"column:payment_category"`
	SnapToken       string `gorm:"column:snap_token"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionNotificationFromMidtrans struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	Amount            string `json:"gross_amount"`
}
