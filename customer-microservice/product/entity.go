package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}