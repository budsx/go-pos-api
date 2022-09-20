package dto

import "go-pos-api/domain"

type LoginResponse struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Merchant int    `json:"merchant"`
	Token    string `json:"token,omitempty"`
}

func UserInput(user domain.User, token string) LoginResponse {
	request := LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Merchant: user.Merchant,
		Token:    token,
	}
	return request
}
