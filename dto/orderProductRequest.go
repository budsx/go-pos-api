package dto

type OrderProductRequest struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	// Price       int    `json:"price"`
}
