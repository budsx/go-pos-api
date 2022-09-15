package dto

type LoginResponse struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Merchant int    `json:"merchant"`
	Token    string `json:"token,omitempty"`
}
