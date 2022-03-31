package main

import (
	"product-microservice/database"
	"product-microservice/product"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// TEST
	router.POST("/api/product/test", func(c *gin.Context) {
		var testInput = struct {
			Message string `json:"message"`
		}{}

		err := c.ShouldBindJSON(testInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
			return
		}

		c.JSON(http.StatusOK, testInput)
		return
	}) // Test and be done by api gateway

	db := database.StartConnection()

	// PRODUCT
	var (
		productRepository = product.NewRepository(db)
		productService    = product.NewService(productRepository)
		productHandler    = product.NewHandler(productService)
	)
	router.GET("/api/product", productHandler.GetAllProductsHandler)                  // Get all products and done by logged in or not logged in customer
	router.GET("/api/product/:category", productHandler.GetProductsByCategoryHandler) // Get products by cateogry and done by logged in or not logged in customer

	router.GET("/api/product/:productUUID", productHandler.GetProductByUUIDHandler) // Get product and done by cart microservice
	router.POST("/api/product/total", productHandler.GetTotalByUUIDHandler)         // Get total of certain products and done by payment microservice

	router.Run(":8081")
}
