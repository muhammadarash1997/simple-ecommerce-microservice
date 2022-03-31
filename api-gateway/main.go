package main

import (
	"api-gateway/auth"
	"api-gateway/cart"
	"api-gateway/database"
	"api-gateway/order"
	"api-gateway/payment"
	"api-gateway/product"
	"api-gateway/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	customerDb := database.StartConnection()

	router.GET("/api/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "api gateway ok"}) }) // Test and be done by api gateway

	// USER
	var (
		userRepository = user.NewRepository(customerDb)
		userService    = user.NewService(userRepository)
		authService    = auth.NewService()
		userHandler    = user.NewHandler(userService, authService)
	)
	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)

	// CART
	var (
		cartHandler = cart.NewHandler()
	)
	router.GET("/api/cart/:userUUID", userHandler.AuthenticateHandler, cartHandler.GetCartByUUIDHandler)            // Get all user cart
	router.POST("/api/cart", userHandler.AuthenticateHandler, cartHandler.AddItemByProductUUIDHandler)              // Add product to cart and be done by logged in customer
	router.POST("/api/cart/update", userHandler.AuthenticateHandler, cartHandler.UpdateQuantityByCartUUIDHandler)   // Update quantity product in cart and be done by logged in customer
	router.DELETE("/api/cart/item/:cartUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByUUIDHandler) // Delete cart and be done by logged in customer

	// PRODUCT
	var (
		productHandler = product.NewHandler()
	)
	router.GET("/api/catalog", productHandler.GetAllProductsHandler)                  // Get all products and done by logged in or not logged in customer
	router.GET("/api/catalog/:category", productHandler.GetProductsByCategoryHandler) // Get products by cateogry and done by logged in or not logged in customer

	// ORDER
	var (
		orderHandler = order.NewHandler()
	)
	router.POST("/api/order/cart/:userUUID", userHandler.AuthenticateHandler, orderHandler.CreateOrderHandler) // Order cart and be done by logged in customer

	// PAYMENT
	var (
		paymentHandler = payment.NewHandler()
	)
	router.POST("/api/order/pay", userHandler.AuthenticateHandler, paymentHandler.CreatePaymentHandler) // Pay order and be done by logged in customer

	router.Run(":8080")
}
