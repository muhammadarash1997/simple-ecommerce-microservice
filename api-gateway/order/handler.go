package order

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

func (this *handler) CreateOrderHandler(c *gin.Context) {
	uuid := c.Params.ByName("userUUID")
	var cartInput []CartInput
	err := c.ShouldBindJSON(&cartInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(cartInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8082/api/order/cart/%s", uuid), bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create order failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderDetailsAdded := []OrderDetail{}
	json.NewDecoder(httpResponse.Body).Decode(&orderDetailsAdded)

	response := helper.APIResponse("Create order success", http.StatusOK, "success", orderDetailsAdded)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}
