package order

type Service interface {
	AddOrderByUUID(id string) (Order, error)
	AddOrderDetailsByOrderUUID(cartInput []CartInput, id string) ([]OrderDetail, error)
}

type service struct {
	orderRepository Repository
}

func NewService(orderRepository Repository) *service {
	return &service{orderRepository}
}

func (this *service) AddOrderByUUID(id string) (Order, error) {
	order := Order{}
	order.UserID = id

	orderAdded, err := this.orderRepository.AddOrder(order)
	if err != nil {
		return orderAdded, err
	}

	return orderAdded, nil
}

func (this *service) AddOrderDetailsByOrderUUID(cartInput []CartInput, id string) ([]OrderDetail, error) {
	orderDetails := []OrderDetail{}

	// Map cartInput to orderDetail
	for _, c := range cartInput {
		orderDetail := OrderDetail{}

		orderDetail.OrderID = id
		orderDetail.ProductID = c.ProductID
		orderDetail.Quantity = c.Quantity

		orderDetails = append(orderDetails, orderDetail)
	}

	orderDetailsAdded, err := this.orderRepository.AddOrderDetails(orderDetails)
	if err != nil {
		return nil, err
	}

	return orderDetailsAdded, nil
}
