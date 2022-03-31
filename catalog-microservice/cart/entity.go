package cart

type Cart struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string `gorm:"type:uuid" json:"user_id"`
	ProductID string `gorm:"type:uuid" json:"product_id"`
	Quantity  uint   `gorm:"type:uint" json:"quantity"`
}
