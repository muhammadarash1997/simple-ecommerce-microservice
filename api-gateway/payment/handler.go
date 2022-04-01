package payment

import (
	"api-gateway/helper"
	"bytes"
	"encoding/json"
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
		Message: "payment microservice ok",
	}

	requestBody, err := json.Marshal(testInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
		return
	}

	// httpRequest, err := http.NewRequest("POST", "http://payment-backend/api/payment/test", bytes.NewBuffer(requestBody))
	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("PAYMENT_MICROSERVICE_URL")+"/api/payment/test", bytes.NewBuffer(requestBody))
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

	// httpRequest, err := http.NewRequest("POST", "http://payment-backend/api/order/pay", bytes.NewBuffer(requestBody))
	httpRequest, err := http.NewRequest("POST", "http://"+os.Getenv("PAYMENT_MICROSERVICE_URL")+"/api/order/pay", bytes.NewBuffer(requestBody))
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
