package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Category string `gorm:"type:varchar(100)" json:"category"`
	Price    int    `gorm:"type:int" json:"price"`
}
