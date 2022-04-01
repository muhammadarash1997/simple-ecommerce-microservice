package database

import (
	"fmt"
	"log"
	"os"
	"payment-microservice/order"
	"payment-microservice/payment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST_PAYMENT")
	dbPort := os.Getenv("DB_PORT_PAYMENT")
	dbUser := os.Getenv("DB_USER_PAYMENT")
	dbPass := os.Getenv("DB_PASS_PAYMENT")
	dbName := os.Getenv("DB_NAME_PAYMENT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to payment database")
		return nil
	}
	fmt.Println("Succes to connect to payment database")

	db.AutoMigrate(&order.Order{})
	db.AutoMigrate(&order.OrderDetail{})
	db.AutoMigrate(&payment.Payment{})

	return db
}
