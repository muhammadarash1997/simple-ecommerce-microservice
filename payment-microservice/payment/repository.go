package payment

import "gorm.io/gorm"

type Repository interface {
	Add(payment Payment) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) Add(payment Payment) (Payment, error) {
	err := this.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
