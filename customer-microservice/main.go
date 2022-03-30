package main

import (
	"customer-microservice/auth"
	"customer-microservice/cart"
	"customer-microservice/database"
	"customer-microservice/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	customerDb := database.StartConnection()

	userRepository := user.NewRepository(customerDb)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := user.NewHandler(userService, authService)

	cartRepository := cart.NewRepository(customerDb)
	cartService := cart.NewService(cartRepository)
	cartHandler := cart.NewHandler(cartService)

	router.GET("/api/cart/:userUUID", cartHandler.GetCartByUUIDHandler)                                              // Get all user cart
	router.POST("/api/cart", userHandler.AuthenticateHandler, cartHandler.AddItemByProductUUIDHandler)               // Add product to cart and be done by logged in customer
	router.POST("/api/cart/:cartUUID", userHandler.AuthenticateHandler, cartHandler.UpdateQuantityByCartUUIDHandler) // Update quantity product in cart and be done by logged in customer
	router.DELETE("/api/cart/item/:cartUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByUUIDHandler)  // Delete cart and be done by logged in customer
	router.DELETE("/api/cart/:userUUID", cartHandler.DeleteUserCartByUUIDHandler)                                    // Delete all cart of user and be done by other microservice

	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)

	router.Run(":8080")
}
