package dto

type ProductRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	MerchantId int    `json:"merchant_id"`
}
