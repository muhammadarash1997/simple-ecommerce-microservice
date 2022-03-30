package main

import (
	"payment-microservice/database"
	"payment-microservice/order"
	"payment-microservice/payment"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.StartConnection()

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository)
	orderHandler := order.NewHandler(orderService)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository, orderRepository)
	paymentHandler := payment.NewHandler(paymentService)

	router.POST("api/order/cart/:userUUID", orderHandler.CreateOrderHandler) // Order
	router.POST("api/order/pay/", paymentHandler.CreatePaymentHandler)       // Pay

	router.Run(":8082")
}
