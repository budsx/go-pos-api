package dto

type DetailOrderRequest struct {
	DetailOrderID int `json:"detail_order_id"`
	ProductID     int `json:"product_id"`
	OrderID       int `json:"order_id"`
	Quantity      int `json:"quantity"`
	SubTotal      int `json:"subtotal"`
}
