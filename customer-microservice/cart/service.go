package cart

type Service interface {
	GetCartByUUID(id string) ([]Cart, error)
	AddItem(cartInput CartInput) (Cart, error)
	UpdateQuantityByCartUUID(id string, quantity uint) (Cart, error)
	DeleteCartByUUID(id string) error
	DeleteUserCartByUUID(id string) error
}

type service struct {
	cartRepository Repository
}

func NewService(cartRepository Repository) *service {
	return &service{cartRepository}
}

func (this *service) GetCartByUUID(id string) ([]Cart, error) {
	cart, err := this.cartRepository.GetAll(id)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (this *service) AddItem(cartInput CartInput) (Cart, error) {
	cart := Cart{}

	cart.UserID = cartInput.UserID
	cart.ProductID = cartInput.ProductID
	cart.Quantity = cartInput.Quantity

	cartAdded, err := this.cartRepository.Add(cart)
	if err != nil {
		return cartAdded, err
	}

	return cartAdded, nil
}

func (this *service) UpdateQuantityByCartUUID(id string, quantity uint) (Cart, error)

func (this *service) DeleteCartByUUID(id string) error {
	err := this.cartRepository.DeleteByUUID(id)
	if err != nil {
		return err
	}

	return nil
}

func (this *service) DeleteUserCartByUUID(id string) error {
	err := this.cartRepository.DeleteAllByUUID(id)
	if err != nil {
		return err
	}

	return nil
}
