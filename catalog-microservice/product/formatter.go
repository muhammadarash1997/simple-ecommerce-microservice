package product

type ProductGottenFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func FormatProductsGotten(allProductsGotten []Product) []ProductGottenFormatter {
	allProductsFormatted := []ProductGottenFormatter{}

	for _, productGotten := range allProductsGotten {
		// Create productFormatted
		productFormatted := ProductGottenFormatter{}

		productFormatted.ID = productGotten.ID
		productFormatted.Name = productGotten.Name
		productFormatted.Category = productGotten.Category
		productFormatted.Price = productGotten.Price

		// Tambahkan productFormatted ke allProductsFormatted
		allProductsFormatted = append(allProductsFormatted, productFormatted)
	}

	return allProductsFormatted
}