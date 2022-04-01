package product

import (
	"net/http"
	"product-microservice/helper"

	"github.com/gin-gonic/gin"
)

type handler struct {
	productService Service
}

func NewHandler(productService Service) *handler {
	return &handler{productService}
}

func (this *handler) GetAllProductsHandler(c *gin.Context) {
	// Call Process
	productsGotten, err := this.productService.GetAllProducts()
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Output
	allProductsFormatted := FormatProductsGotten(productsGotten)
	c.JSON(http.StatusOK, allProductsFormatted)
	return
}

func (this *handler) GetProductsByCategoryHandler(c *gin.Context) {
	// Get path params
	categoryRequest := c.Params.ByName("category")

	// Call Process
	productsGotten, err := this.productService.GetProductByCategory(categoryRequest)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Output
	allProductsFormatted := FormatProductsGotten(productsGotten)
	c.JSON(http.StatusOK, allProductsFormatted)
}

func (this *handler) GetProductByUUIDHandler(c *gin.Context) {
	// Get path params
	uuid := c.Params.ByName("productUUID")

	// Call process
	product, err := this.productService.GetProductByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	// Output
	c.JSON(http.StatusOK, product)
	return
}

func (this *handler) GetTotalByUUIDHandler(c *gin.Context) {
	// Read payload
	var requestObjects []RequestModel
	err := c.ShouldBindJSON(&requestObjects)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get total failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Call process
	total, err := this.productService.GetTotal(requestObjects)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get total failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Output
	c.JSON(http.StatusOK, gin.H{"total": total})
	return
}
