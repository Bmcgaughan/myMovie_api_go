package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck is a handler for the healthcheck endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

