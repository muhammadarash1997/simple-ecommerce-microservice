package cart

import "customer-microservice/product"

type CartGottenFormatted struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func FormatProductGotten(product product.Product, quantity int) CartGottenFormatted {
	cartGottenFormatted := CartGottenFormatted{}

	cartGottenFormatted.ID = product.ID
	cartGottenFormatted.Name = product.Name
	cartGottenFormatted.Category = product.Category
	cartGottenFormatted.Price = product.Price
	cartGottenFormatted.Quantity = quantity

	return cartGottenFormatted
}
