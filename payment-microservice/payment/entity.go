package payment

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID      string `gorm:"primaryKey;type:uuid;default:get_random_uuid()" json:"id"`
	OrderID string `gorm:"type:uuid" json:"order_id"`
	UserID  string `gorm:"type:uuid" json:"user_id"`
	Total   int    `gorm:"type:int" json:"total"`
	Status  bool   `gorm:"type:boolean" json:"status"`
}
