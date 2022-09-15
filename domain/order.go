package domain

import "time"

type Order struct {
	OrderID      int       `gorm:"column:order_id"`
	UserID       int       `gorm:"column:user_id"`
	CustomerName string    `gorm:"column:customer_name"`
	Amount       int       `gorm:"column:amount"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
