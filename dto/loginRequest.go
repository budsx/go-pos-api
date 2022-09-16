package dto

import "go-pos-api/domain"

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
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
