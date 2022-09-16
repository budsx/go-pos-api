package domain

type User struct {
	ID       int    `json:"user_id" gorm:"column:user_id; primary_key:auto_increment"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Role     int    `json:"role" gorm:"column:role"`
	Merchant int    `json:"merchant_id" gorm:"column:merchant_id"`
}
