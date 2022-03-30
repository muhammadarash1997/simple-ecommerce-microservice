package user

import (
	"customer-microservice/auth"
	"customer-microservice/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService Service
	authService auth.Service
}

func NewHandler(userService Service, authService auth.Service) *handler {
	return &handler{userService, authService}
}

func (this *handler) RegisterUserHandler(c *gin.Context) {
	// catch input from user
	// mapping input from user to RegisterUserInput struct
	// passing that struct as service parameter

	var userInput RegisterUserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userRegistered, err := this.userService.RegisterUser(userInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userFormatted := FormatUserRegistered(userRegistered)

	response := helper.APIResponse("Register account succes", http.StatusOK, "success", userFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *handler) LoginHandler(c *gin.Context) {
	var loginInput LoginInput

	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLogged, err := this.userService.Login(loginInput)
	if err != nil {
		errormessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	tokenGenerated, err := this.authService.GenerateToken(userLogged.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormatted := FormatUserLogged(userLogged, tokenGenerated)

	response := helper.APIResponse("Login success", http.StatusOK, "success", userFormatted)

	c.JSON(http.StatusOK, response)
}

func (this *handler) AuthenticateHandler(c *gin.Context) {
	// Ambil token dari header
	tokenInput := c.GetHeader("Authorization")

	// Validasi apakah benar itu adalah bearer token
	if !strings.Contains(tokenInput, "Bearer") {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	tokenWithoutBearer := strings.Split(tokenInput, " ")[1]

	// Validasi token apakah berlaku
	token, err := this.authService.ValidateToken(tokenWithoutBearer)
	if err != nil {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// Mengubah token yang bertipe jwt.Token menjadi bertipe jwt.MapClaims
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id := claim["user_uuid"].(string)
	user, err := this.userService.GetUser(id)
	if err != nil {
		response := helper.APIResponse("Invalid token", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", user)
}
