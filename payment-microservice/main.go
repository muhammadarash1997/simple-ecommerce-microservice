package main

import (
	"net/http"
	"payment-microservice/database"
	"payment-microservice/order"
	"payment-microservice/payment"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// TEST
	router.GET("/api/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "payment microservice ok"}) }) // Test and be done by api gateway

	db := database.StartConnection()

	// ORDER
	var (
		orderRepository = order.NewRepository(db)
		orderService    = order.NewService(orderRepository)
		orderHandler    = order.NewHandler(orderService)
	)
	router.POST("/api/order/cart/:userUUID", orderHandler.CreateOrderHandler) // Order cart and be done by logged in customer

	// PAYMENT
	var (
		paymentRepository = payment.NewRepository(db)
		paymentService    = payment.NewService(paymentRepository, orderRepository)
		paymentHandler    = payment.NewHandler(paymentService)
	)
	router.POST("/api/order/pay", paymentHandler.CreatePaymentHandler) // Pay order and be done by logged in customer

	router.Run(":8082")
}
