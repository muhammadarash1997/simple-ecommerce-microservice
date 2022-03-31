package product

type ProductGottenFormatted struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func FormatProductsGotten(allProductsGotten []Product) []ProductGottenFormatted {
	allProductsFormatted := []ProductGottenFormatted{}

	for _, productGotten := range allProductsGotten {
		// Create productFormatted
		productFormatted := ProductGottenFormatted{}

		productFormatted.ID = productGotten.ID
		productFormatted.Name = productGotten.Name
		productFormatted.Category = productGotten.Category
		productFormatted.Price = productGotten.Price

		// Tambahkan productFormatted ke allProductsFormatted
		allProductsFormatted = append(allProductsFormatted, productFormatted)
	}

	return allProductsFormatted
}
