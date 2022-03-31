package cart

type Cart struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string `gorm:"type:uuid" json:"user_id"`
	ProductID string `gorm:"type:uuid" json:"product_id"`
	Quantity  uint   `gorm:"type:uint" json:"quantity"`
}

type CartGottenFormatted struct {
	Product Product `json:"product"`
	Quantity uint `json:"quantity"`
}

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type UpdateQuantityInput struct {
	CartID string `json:"cart_id"`
	Quantity uint   `json:"quantity"`
}