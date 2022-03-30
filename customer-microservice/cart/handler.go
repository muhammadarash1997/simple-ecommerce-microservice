package cart

import (
	"customer-microservice/helper"
	"customer-microservice/product"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cartService Service
}

func NewHandler(cartService Service) *handler {
	return &handler{cartService}
}

func (this *handler) GetCartByUUIDHandler(c *gin.Context) {
	// Read param
	// Call process
	// Get detail product
	// Output

	// Read param
	uuid := c.Params.ByName("userUUID")

	// Call process
	cartGotten, err := this.cartService.GetCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get detail product
	cartResponse := []CartGottenFormatted{}

	for _, content := range cartGotten {
		client := &http.Client{}

		request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/api/catalog/%s", content.ProductID), nil)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response, err := client.Do(request)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		product := product.Product{}
		json.NewDecoder(response.Body).Decode(&product)
		cartFormatted := FormatProductGotten(product, content.Quantity)
		cartResponse = append(cartResponse, cartFormatted)

		response.Body.Close()
	}

	// Output
	response := helper.APIResponse("Get cart success", http.StatusOK, "success", cartResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (this *handler) AddItemByProductUUIDHandler(c *gin.Context) {
	// Read param
	// Call process
	// Output

	// Read param
	var cartInput CartInput

	err := c.ShouldBindJSON(&cartInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call process
	cartAdded, err := this.cartService.AddItem(cartInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Output
	response := helper.APIResponse("Add cart success", http.StatusOK, "success", cartAdded)
	c.JSON(http.StatusOK, response)
	return
}

func (this *handler) UpdateQuantityByCartUUIDHandler(c *gin.Context)

func (this *handler) DeleteCartByUUIDHandler(c *gin.Context) {
	// Read param
	uuid := c.Params.ByName("cartUUID")

	// Call process
	err := this.cartService.DeleteCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Output
	response := helper.APIResponse("Delete cart success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}

func (this *handler) DeleteUserCartByUUIDHandler(c *gin.Context) {
	// Read param
	uuid := c.Params.ByName("userUUID")

	// Call process
	err := this.cartService.DeleteUserCartByUUID(uuid)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete user cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Output
	c.JSON(http.StatusOK, nil)
	return
}
