package cart

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string `gorm:"type:uuid" json:"user_id"`
	ProductID string `gorm:"type:uuid" json:"product_id"`
	Quantity  int    `gorm:"type:int" json:"quantity"`
}
