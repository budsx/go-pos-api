package domain

type DetailOrder struct {
	DetailOrderID int `gorm:"column:detail_order_id; primary_key:auto_increment"`
	ProductID     int `gorm:"column:product_id"`
	OrderID       int `gorm:"column:order_id"`
	Quantity      int `gorm:"column:quantity"`
	SubTotal      int `gorm:"column:subtotal"`
}

func (DetailOrder) TableName() string {
	return "detail_order"
}
