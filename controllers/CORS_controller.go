package controllers

import (
	"context"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		// Add CORS headers
		c.Header("Access-Control-Allow-Origin", "http://example.com")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

		// Prepare response
		c.JSON(http.StatusOK, gin.H{
			"message": "this response has custom headers",
		})
	}
}
