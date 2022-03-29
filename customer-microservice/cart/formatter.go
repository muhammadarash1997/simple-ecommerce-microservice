package cart

import "customer-microservice/product"

type CartGottenFormatted struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Price     int    `json:"price"`
	Quantity  uint   `json:"quantity"`
}

func FormatProductGotten(product product.Product, quantity uint) CartGottenFormatted {
	cartGottenFormatted := CartGottenFormatted{}

	cartGottenFormatted.ProductID = product.ID
	cartGottenFormatted.Name = product.Name
	cartGottenFormatted.Category = product.Category
	cartGottenFormatted.Price = product.Price
	cartGottenFormatted.Quantity = quantity

	return cartGottenFormatted
}
