package order

type CartInput struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Price     int    `json:"price"`
	Quantity  uint   `json:"quantity"`
}
