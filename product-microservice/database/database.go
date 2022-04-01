package database

import (
	"fmt"
	"log"
	"os"
	"product-microservice/product"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST_PRODUCT")
	dbPort := os.Getenv("DB_PORT_PRODUCT")
	dbUser := os.Getenv("DB_USER_PRODUCT")
	dbPass := os.Getenv("DB_PASS_PRODUCT")
	dbName := os.Getenv("DB_NAME_PRODUCT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to product database")
		return nil
	}
	fmt.Println("Succes to connect to product database")

	db.AutoMigrate(&product.Product{})

	return db
}
