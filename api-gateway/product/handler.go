package product

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

func (h *handler) Test(c *gin.Context) {
	client := &http.Client{}

	var testInput = struct {
		Message string `json:"message"`
	}{
		Message: "product microservice ok",
	}

	requestBody, err := json.Marshal(testInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	httpRequest, err := http.NewRequest("POST", "http://localhost:8081/api/product/test", bytes.NewBuffer(requestBody))
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

func (this *handler) GetAllProductsHandler(c *gin.Context) {
	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", "http://localhost:8081/api/catalog", nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get products failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productResponse := []ProductGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&productResponse)

	response := helper.APIResponse("Get products success", http.StatusOK, "success", productResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}

func (this *handler) GetProductsByCategoryHandler(c *gin.Context) {
	categoryRequest := c.Params.ByName("category")

	client := &http.Client{}

	httpRequest, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/api/cart/%s", categoryRequest), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get products by category failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Get products by category failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productResponse := []ProductGottenFormatted{}
	json.NewDecoder(httpResponse.Body).Decode(&productResponse)

	response := helper.APIResponse("Get roducts by category success", http.StatusOK, "success", productResponse)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}
