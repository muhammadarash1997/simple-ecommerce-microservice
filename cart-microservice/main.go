package main

import (
	"cart-microservice/cart"
	"cart-microservice/database"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// TEST
	router.GET("/api/test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "cart microservice ok"}) }) // Test and be done by api gateway

	db := database.StartConnection()

	// CART
	var (
		cartRepository = cart.NewRepository(db)
		cartService    = cart.NewService(cartRepository)
		cartHandler    = cart.NewHandler(cartService)
	)
	router.GET("/api/cart/:userUUID", cartHandler.GetCartByUUIDHandler)            // Get all user cart and be done by logged in customer
	router.POST("/api/cart", cartHandler.AddItemByProductUUIDHandler)              // Add product to cart and be done by logged in customer
	router.POST("/api/cart/update", cartHandler.UpdateQuantityByCartUUIDHandler)   // Update quantity product in cart and be done by logged in customer
	router.DELETE("/api/cart/item/:cartUUID", cartHandler.DeleteCartByUUIDHandler) // Delete cart and be done by logged in customer

	router.DELETE("/api/cart/:userUUID", cartHandler.DeleteUserCartByUUIDHandler) // Delete all cart of user and be done by other payment microservice

	router.Run(":8083")
}
