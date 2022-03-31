package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"payment-microservice/order"
)

type Service interface {
	AddPayment(orderInput OrderInput) (Payment, error)
}

type service struct {
	paymentRepository Repository
	orderRepository   order.Repository
}

func NewService(paymentRepository Repository, orderRepository order.Repository) *service {
	return &service{paymentRepository, orderRepository}
}

func (this *service) AddPayment(orderInput OrderInput) (Payment, error) {
	payment := Payment{}

	payment.OrderID = orderInput.ID
	payment.UserID = orderInput.UserID

	// Get ProductID-ProductID and Quantity di OrderDetail berdasarkan OrderID <-- Gampangnya ambil semua OrderDetail
	orderDetails, err := this.orderRepository.GetOrderDetails(orderInput.ID)
	if err != nil {
		return payment, err
	}

	// Get Price from Product table dengan passing quantity
	type requestModel struct {
		ProductID string `json:"product_id"`
		Quantity  uint   `json:"quantity"`
	}

	client := &http.Client{}

	var requestObjects []requestModel
	for _, c := range orderDetails {
		requestObject := requestModel{
			ProductID: c.ProductID,
			Quantity:  c.Quantity,
		}

		requestObjects = append(requestObjects, requestObject)
	}

	requestBody, err := json.Marshal(requestObjects)
	if err != nil {
		return payment, err
	}

	httpRequest, err := http.NewRequest("POST", "http://localhost:8081/api/product/total", bytes.NewBuffer(requestBody))
	if err != nil {
		return payment, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	response, err := client.Do(httpRequest)
	if err != nil {
		return payment, err
	}
	defer response.Body.Close()

	var Data = struct {
		Total int `json:"total"`
	}{}
	err = json.NewDecoder(response.Body).Decode(&Data)
	if err != nil {
		return payment, err
	}

	payment.Total = Data.Total

	paymentAdded, err := this.paymentRepository.Add(payment)
	if err != nil {
		return paymentAdded, err
	}

	return payment, nil
}
