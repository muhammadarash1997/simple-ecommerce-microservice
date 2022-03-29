package product

import (
	"catalog-microservice/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	productService Service
}

func NewHandler(productService Service) *handler {
	return &handler{productService}
}

func (this *handler) GetAllProductsHandler(c *gin.Context) {
	// Call process
	// Output

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
	response := helper.APIResponse("Get products successfully", http.StatusOK, "success", allProductsFormatted)
	c.JSON(http.StatusOK, response)
}

func (this *handler) GetProductsByCategoryHandler(c *gin.Context) {
	// Get path params
	// Call process
	// Output

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
	response := helper.APIResponse("Get products successfully", http.StatusOK, "success", allProductsFormatted)
	c.JSON(http.StatusOK, response)
}

func (this *handler) GetProductByUUIDHandler(c *gin.Context) {
	// Get path params
	// Call process
	// Output

	// Get path params
	uuid := c.Params.ByName("productUUID")

	// Call process
	product, err := this.productService.GetProductByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"message": "Invalid query"}
		response := helper.APIResponse("Get product failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Output
	c.JSON(http.StatusOK, product)
}
