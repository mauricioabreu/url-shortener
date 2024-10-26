package serializers

import "github.com/gin-gonic/gin"

type StandardResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	response := StandardResponse{
		Success: false,
		Error:   message,
	}

	c.JSON(code, response)
}

func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	response := StandardResponse{
		Success: true,
		Data:    data,
	}

	c.JSON(code, response)
}
