package dto

import "go-pos-api/domain"

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     int    `json:"role"`
	Merchant int    `json:"merchant"`
	Token    string `json:"-"`
}

func UserInput(user domain.User, token string) UserDTO {
	register := UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Merchant: user.Merchant,
		Token:    token,
	}
	return register
}
