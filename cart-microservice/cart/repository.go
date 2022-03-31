package cart

import "gorm.io/gorm"

type Repository interface {
	GetAll(id string) ([]Cart, error)
	Add(cart Cart) (Cart, error)
	UpdateQuantity(id string, quantity uint) error
	DeleteByUUID(id string) error
	DeleteAllByUUID(id string) error
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

func (this *repository) UpdateQuantity(id string, quantity uint) error {
	err := this.db.Model(&Cart{}).Where("id = ?", id).Update("quantity", quantity).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) DeleteByUUID(id string) error {
	err := this.db.Delete("WHERE id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (this *repository) DeleteAllByUUID(id string) error {
	err := this.db.Delete("WHERE user_id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
