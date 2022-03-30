package main

import (
	"catalog-microservice/database"
	"catalog-microservice/product"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.StartConnection()

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := product.NewHandler(productService)

	router.GET("/api/catalog", productHandler.GetAllProductsHandler)                  // Get all products and done by logged in or not logged in customer
	router.GET("/api/catalog/:category", productHandler.GetProductsByCategoryHandler) // Get products by cateogry and done by logged in or not logged in customer

	router.GET("/api/catalog/:productUUID", productHandler.GetProductByUUIDHandler) // Get product and done by microservice
	router.POST("/api/catalog/total", productHandler.GetTotalByUUIDHandler)         // Get total of certain products and done by microservice

	router.Run(":8081")
}
