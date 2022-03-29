package product

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Product, error)
	GetByCategory(category string) ([]Product, error)
	GetByUUID(uuid string) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (this *repository) GetAll() ([]Product, error) {
	products := []Product{}
	err := this.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (this *repository) GetByCategory(category string) ([]Product, error) {
	products := []Product{}
	err := this.db.Find(&products).Where("category = ?", category).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (this *repository) GetByUUID(uuid string) (Product, error) {
	product := Product{}
	err := this.db.Find(&product).Where("id = ?", uuid).Error
	if err != nil {
		return product, err
	}

	return product, nil
}