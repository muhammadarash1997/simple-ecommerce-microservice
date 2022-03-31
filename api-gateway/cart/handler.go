package cart

import (
	"api-gateway/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (this *handler) GetCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")

	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/api/cart/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartResponse := []CartGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&cartResponse)

	response := helper.APIResponse("Get cart success", http.StatusOK, "success", cartResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) AddItemByProductUUIDHandler(c *gin.Context) {
	var cartInput CartInput

	err := c.ShouldBindJSON(&cartInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(cartInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpRequest, err := http.NewRequest("POST", "http://localhost:8081/api/cart", bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartResponse := Cart{}
	json.NewDecoder(httpResponse.Body).Decode(&cartResponse)

	response := helper.APIResponse("Add cart success", http.StatusOK, "success", cartResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) UpdateQuantityByCartUUIDHandler(c *gin.Context) {
	var updateQuantityInput UpdateQuantityInput

	err := c.ShouldBindJSON(&updateQuantityInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(updateQuantityInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Update quantity failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpRequest, err := http.NewRequest("POST", "http://localhost:8081/api/cart/update", bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Update quantity failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Update quantity failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update quantity success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) DeleteCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("cartUUID")

	client := &http.Client{}

	httpRequest, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8081/api/cart/item/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete cart success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}