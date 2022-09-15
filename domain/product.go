package domain

type Product struct {
	ID         int    `gorm:"column:product_id"`
	Name       string `gorm:"column:name"`
	Price      int    `gorm:"column:price"`
	Stock      int    `gorm:"column:stock"`
	MerchantId int    `gorm:"column:merchant_id"`
}
