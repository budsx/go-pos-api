package domain

import "time"

type Order struct {
	OrderID      int
	UserID       int
	CustomerName string
	Amount       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
