package database

import (
	"api-gateway/user"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartConnection() *gorm.DB {
	dbHost := os.Getenv("DB_HOST_APIGATEWAY")
	dbPort := os.Getenv("DB_PORT_APIGATEWAY")
	dbUser := os.Getenv("DB_USER_APIGATEWAY")
	dbPass := os.Getenv("DB_PASS_APIGATEWAY")
	dbName := os.Getenv("DB_NAME_APIGATEWAY")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		fmt.Println("Failed to connect to customer database")
		return nil
	}
	fmt.Println("Succes to connect to customer database")

	db.AutoMigrate(&user.User{})

	return db
}
