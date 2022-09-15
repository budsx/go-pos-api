package dto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Role     int    `json:"role" form:"role" binding:"required"`
	Merchant int    `json:"merchant_id" form:"merchant" binding:"required"`
}
