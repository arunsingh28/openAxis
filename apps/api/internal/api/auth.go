package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the credentials are valid
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/login", Login)
}
