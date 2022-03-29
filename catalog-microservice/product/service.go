package product

type Service interface {
	GetAllProducts() ([]Product, error)
	GetProductByCategory(category string) ([]Product, error)
	GetProductByUUID(uuid string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (this *service) GetAllProducts() ([]Product, error) {
	products, err := this.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (this *service) GetProductByCategory(category string) ([]Product, error) {
	products, err := this.repository.GetByCategory(category)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (this *service) GetProductByUUID(uuid string) (Product, error) {
	product, err := this.repository.GetByUUID(uuid)
	if err != nil {
		return product, err
	}

	return product, nil
}