package handler

import (
	"net/http"

	"todo-app/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var credentials service.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := handler.AuthService.Authenticate(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "id": user.ID})
}
