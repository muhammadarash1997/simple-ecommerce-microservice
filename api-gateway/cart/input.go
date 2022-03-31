package cart

type CartInput struct {
	UserID    string `json:"user_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Quantity  uint   `json:"quantity" binding:"required"`
}
