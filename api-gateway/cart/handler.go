package cart

import (
	"api-gateway/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Test(c *gin.Context) {
	client := &http.Client{}

	var testInput = struct {
		Message string `json:"message"`
	}{
		Message: "cart microservice ok",
	}

	requestBody, err := json.Marshal(testInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	// httpRequest, err := http.NewRequest("POST", "http://cart-backend/api/cart/test", bytes.NewBuffer(requestBody))
	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("CART_MICROSERVICE_URL")+"/api/cart/test", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	var testOutput = struct {
		Message string `json:"message"`
	}{}

	json.NewDecoder(httpResponse.Body).Decode(&testOutput)
	c.JSON(http.StatusOK, testOutput)

	httpResponse.Body.Close()
	return
}

func (this *handler) GetCartByUUIDHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")

	client := &http.Client{}

	// httpRequest, err := http.NewRequest("GET", fmt.Sprintf("http://cart-backend/api/cart/%s", uuid), nil)
	httpRequest, err := http.NewRequest("GET", "http://"+os.Getenv("CART_MICROSERVICE_URL")+fmt.Sprintf("/api/cart/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartGottenFormatted := []CartGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&cartGottenFormatted)

	response := helper.APIResponse("Get cart success", http.StatusOK, "success", cartGottenFormatted)
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
		response := helper.APIResponse("Add cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// httpRequest, err := http.NewRequest("POST", "http://cart-backend/api/cart", bytes.NewBuffer(requestBody))
	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("CART_MICROSERVICE_URL")+"/api/cart", bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Add cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartAdded := Cart{}
	json.NewDecoder(httpResponse.Body).Decode(&cartAdded)

	response := helper.APIResponse("Add cart success", http.StatusOK, "success", cartAdded)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) UpdateQuantityByCartUUIDHandler(c *gin.Context) {
	var updateQuantityInput UpdateQuantityInput

	err := c.ShouldBindJSON(&updateQuantityInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update quantity failed", http.StatusUnprocessableEntity, "error", errorMessage)
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

	// httpRequest, err := http.NewRequest("POST", "http://cart-backend/api/cart/update", bytes.NewBuffer(requestBody))
	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("CART_MICROSERVICE_URL")+"/api/cart/update", bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Update quantity failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

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

	// httpRequest, err := http.NewRequest("DELETE", fmt.Sprintf("http://cart-backend/api/cart/item/%s", uuid), nil)
	httpRequest, err := http.NewRequest("DELETE", "http://"+os.Getenv("CART_MICROSERVICE_URL")+fmt.Sprintf("/api/cart/item/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

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
