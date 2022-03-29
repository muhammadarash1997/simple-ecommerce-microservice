package order

import "gorm.io/gorm"

type Repository interface {
	AddOrder(order Order) (Order, error)
	AddOrderDetails(order []OrderDetail) ([]OrderDetail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) AddOrder(order Order) (Order, error) {
	err := this.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (this *repository) AddOrderDetails(orderDetails []OrderDetail) ([]OrderDetail, error) {
	err := this.db.Create(&orderDetails).Error
	if err != nil {
		return orderDetails, err
	}

	return orderDetails, nil
}