package product

type RequestModel struct {
	ProductID string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}
