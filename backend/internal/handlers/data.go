package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple API Handler

func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Go backend",
	})
}
