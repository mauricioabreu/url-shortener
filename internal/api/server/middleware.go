package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutResponse(c *gin.Context) {
	c.JSON(
		http.StatusRequestTimeout,
		gin.H{"error": "A timeout error occurred. We are working to solve it as soon as possible"},
	)
}

func timeoutMiddleware(srvTimeout time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(srvTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}