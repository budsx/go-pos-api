package dto

import "go-pos-api/domain"

type RegisterResponse struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Merchant int    `json:"merchant_id"`
}

func RegisterInput(user domain.User, token string) RegisterResponse {
	response := RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Merchant: user.Merchant,
	}
	return response
}
