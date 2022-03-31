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
	db := database.StartConnection()

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := order.NewHandler(orderService)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository, orderRepository)
	paymentHandler := payment.NewHandler(paymentService)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) }) // Test

	router.POST("/api/order/cart/:userUUID", orderHandler.CreateOrderHandler) // Order
	router.POST("/api/order/pay/", paymentHandler.CreatePaymentHandler)       // Pay

	router.Run(":8082")
}
