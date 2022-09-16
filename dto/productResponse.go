package dto

type ProductResponse struct {
	ID         int    `json:"product_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	MerchantId int    `json:"merchant_id"`
}
