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

	router.GET("/api/cart/:userUUID", cartHandler.GetCartByUUIDHandler)
	router.POST("/api/cart", userHandler.AuthenticateHandler, cartHandler.AddItemByProductUUIDHandler)
	router.POST("/api/cart/:cartUUID", userHandler.AuthenticateHandler, cartHandler.UpdateQuantityByCartUUIDHandler)  // now
	router.DELETE("/api/cart/:cartUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByCartUUIDHandler)
	router.DELETE("/api/cart/:userUUID", userHandler.AuthenticateHandler, cartHandler.DeleteCartByUserUUIDHandler)

	router.POST("/api/user/register", userHandler.RegisterUserHandler)
	router.POST("/api/user/login", userHandler.LoginHandler)

	router.Run(":8080")
}
