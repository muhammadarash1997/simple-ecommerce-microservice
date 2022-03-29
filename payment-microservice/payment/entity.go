package payment

type Payment struct {
	ID string `json:"id"`
	OrderID string `json:"order_id"`
	UserID string `json:"user_id"`
}