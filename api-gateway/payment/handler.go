package payment

import (
	"api-gateway/helper"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (this *handler) CreatePaymentHandler(c *gin.Context) {
	var orderInput OrderInput
	err := c.ShouldBindJSON(&orderInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create payment failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	client := &http.Client{}

	requestBody, err := json.Marshal(orderInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create payment failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	httpRequest, err := http.NewRequest("POST", "http://localhost:8082/api/order/pay", bytes.NewBuffer(requestBody))
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create payment failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Create payment failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	paymentAdded := Payment{}
	json.NewDecoder(httpResponse.Body).Decode(&paymentAdded)

	response := helper.APIResponse("Create payment success", http.StatusOK, "success", paymentAdded)
	c.JSON(http.StatusOK, response)

	httpResponse.Body.Close()
	return
}
