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
	router.POST("/api/cart/test", func(c *gin.Context) {
		var testInput = struct {
			Message string `json:"message"`
		}{}

		err := c.ShouldBindJSON(&testInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
			return
		}

		c.JSON(http.StatusOK, testInput)
		return
	}) // Test and be done by api gateway

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

	router.Run(":8080") // Karena microservice ini akan dijadikan container maka 8080 di sini adalah localhost:8080 nya si container bukan localhost:8080 nya si host
}
