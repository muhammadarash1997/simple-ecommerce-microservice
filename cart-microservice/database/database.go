package database

import (
	"cart-microservice/cart"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST_CART")
	dbPort := os.Getenv("DB_PORT_CART")
	dbUser := os.Getenv("DB_USER_CART")
	dbPass := os.Getenv("DB_PASS_CART")
	dbName := os.Getenv("DB_NAME_CART")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to cart database")
		return nil
	}
	fmt.Println("Succes to connect to cart database")

	db.AutoMigrate(&cart.Cart{})

	return db
}
