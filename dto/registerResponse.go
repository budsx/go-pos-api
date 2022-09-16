package dto

type RegisterResponse struct {
	ID       int    `json:"user_id" `
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Merchant int    `json:"merchant_id"`
}
