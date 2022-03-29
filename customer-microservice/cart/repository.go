package cart

import "gorm.io/gorm"

type Repository interface {
	GetAll(id string) ([]Cart, error)
	Add(cart Cart) (Cart, error)
	UpdateQuantity(id string, quantity uint) (Cart, error)
	DeleteByUUID(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) GetAll(id string) ([]Cart, error) {
	cart := []Cart{}
	err := this.db.Where("user_id == ?", id).Find(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (this *repository) Add(cart Cart) (Cart, error) {
	err := this.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (this *repository) UpdateQuantity(id string, quantity uint) (Cart, error)

func (this *repository) DeleteByUUID(id string) error