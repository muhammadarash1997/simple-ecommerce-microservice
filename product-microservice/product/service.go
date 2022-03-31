package product

type Service interface {
	GetAllProducts() ([]Product, error)
	GetProductByCategory(category string) ([]Product, error)
	GetProductByUUID(uuid string) (Product, error)
	GetTotal(requestObjects []RequestModel) (int, error)
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

func (this *service) GetTotal(requestObjects []RequestModel) (int, error) {
	type total struct {
		Price    int
		Quantity uint
	}

	var totalObjects []total

	for _, c := range requestObjects {
		product, err := this.repository.GetByUUID(c.ProductID)
		if err != nil {
			return -1, err
		}

		totalObject := total{}
		totalObject.Price = product.Price
		totalObject.Quantity = c.Quantity

		totalObjects = append(totalObjects, totalObject)
	}

	var totalPayment int
	for _, c := range totalObjects {
		totalPayment += int(c.Quantity) * c.Price
	}

	return totalPayment, nil
}
