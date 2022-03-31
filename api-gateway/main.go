package main

import (
	"api-gateway/auth"
	"api-gateway/cart"
	"api-gateway/database"
	"api-gateway/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	customerDb := database.StartConnection()

	/// USER ///
	var (
		userRepository = user.NewRepository(customerDb)
		userService = user.NewService(userRepository)
		authService = auth.NewService()
		userHandler = user.NewHandler(userService, authService)
	)
	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)

	/// CART ///
	var (
		cartHandler = cart.NewHandler()
	)
	router.GET("/api/cart/:userUUID", userHandler.AuthenticateHandler, cartHandler.GetCartByUUIDHandler)            // Get all user cart
	router.POST("/api/cart", userHandler.AuthenticateHandler, cartHandler.AddItemByProductUUIDHandler)              // Add product to cart and be done by logged in customer
	router.POST("/api/cart/update", userHandler.AuthenticateHandler, cartHandler.UpdateQuantityByCartUUIDHandler)   // Update quantity product in cart and be done by logged in customer
	router.DELETE("/api/cart/item/:cartUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByUUIDHandler) // Delete cart and be done by logged in customer

	router.Run(":8080")
}
