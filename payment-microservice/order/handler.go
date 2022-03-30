package order

import (
	"fmt"
	"net/http"
	"payment-microservice/helper"

	"github.com/gin-gonic/gin"
)

type handler struct {
	orderService Service
}

func NewHandler(orderService Service) *handler {
	return &handler{orderService}
}

func (this *handler) CreateOrderHandler(c *gin.Context) {
	// Read payload
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

	// Create Order
	orderAdded, err := this.orderService.AddOrderByUUID(uuid)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Create OrderDetails
	orderDetailsAdded, err := this.orderService.AddOrderDetailsByOrderUUID(cartInput, orderAdded.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create order details failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Delete cart
	client := &http.Client{}

	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/api/cart/%s", uuid), nil)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = client.Do(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Delete cart failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Output
	response := helper.APIResponse("Create order successfully", http.StatusOK, "success", orderDetailsAdded)
	c.JSON(http.StatusOK, response)
	return
}
