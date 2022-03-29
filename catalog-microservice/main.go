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

	router.GET("/api/catalog", productHandler.GetAllProductsHandler)
	router.GET("/api/catalog/:category", productHandler.GetProductsByCategoryHandler)

	// For other microservice
	router.GET("/api/catalog/:productUUID", productHandler.GetProductByUUIDHandler)

	router.Run(":8081")
}
