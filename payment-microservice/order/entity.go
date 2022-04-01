package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID string `gorm:"type:uuid" json:"user_id"`
}

type OrderDetail struct {
	gorm.Model
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	OrderID   string `gorm:"type:uuid" json:"order_id"`
	ProductID string `gorm:"type:uuid" json:"product_id"`
	Quantity  uint   `gorm:"type:uint" json:"quantity"`
}
