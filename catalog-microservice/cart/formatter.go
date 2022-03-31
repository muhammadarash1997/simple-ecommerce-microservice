package cart

import "catalog-microservice/product"

type CartGottenFormatted struct {
	Product product.Product `json:"product"`
	Quantity uint `json:"quantity"`
}

func FormatProductGotten(product product.Product, quantity uint) CartGottenFormatted {
	cartGottenFormatted := CartGottenFormatted{}

	cartGottenFormatted.Product = product
	cartGottenFormatted.Quantity = quantity

	return cartGottenFormatted
}