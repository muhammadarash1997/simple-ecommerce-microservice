package main

import (
	"payment-microservice/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.StartConnection()

	router.POST("api/order/cart/:userUUID", orderHandler.CreateOrderHandler)

	router.Run(":8082")
}
