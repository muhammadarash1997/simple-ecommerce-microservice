package main

import (
	"catalog-microservice/cart"
	"catalog-microservice/database"
	"catalog-microservice/product"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.StartConnection()

	/// PRODUCT ///
	var (
		productRepository = product.NewRepository(db)
		productService    = product.NewService(productRepository)
		productHandler    = product.NewHandler(productService)
	)
	router.GET("/api/catalog", productHandler.GetAllProductsHandler)                  // Get all products and done by logged in or not logged in customer
	router.GET("/api/catalog/:category", productHandler.GetProductsByCategoryHandler) // Get products by cateogry and done by logged in or not logged in customer

	router.GET("/api/catalog/:productUUID", productHandler.GetProductByUUIDHandler) // Get product and done by microservice
	router.POST("/api/catalog/total", productHandler.GetTotalByUUIDHandler)         // Get total of certain products and done by microservice

	/// CART ///
	var (
		cartRepository = cart.NewRepository(db)
		cartService    = cart.NewService(cartRepository)
		cartHandler    = cart.NewHandler(cartService)
	)
	router.GET("/api/cart/:userUUID", cartHandler.GetCartByUUIDHandler)            // Get all user cart and be done by logged in customer
	router.POST("/api/cart", cartHandler.AddItemByProductUUIDHandler)              // Add product to cart and be done by logged in customer
	router.POST("/api/cart/update", cartHandler.UpdateQuantityByCartUUIDHandler)   // Update quantity product in cart and be done by logged in customer
	router.DELETE("/api/cart/item/:cartUUID", cartHandler.DeleteCartByUUIDHandler) // Delete cart and be done by logged in customer

	router.DELETE("/api/cart/:userUUID", cartHandler.DeleteUserCartByUUIDHandler) // Delete all cart of user and be done by other microservice

	router.Run(":8081")
}
