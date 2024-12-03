package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int         `json:"status"`          // HTTP Status Code
	Message string      `json:"message"`         // Response message
	Data    interface{} `json:"data,omitempty"`  // Success data (optional)
	Error   interface{} `json:"error,omitempty"` // Error details (optional)
}

func RespondJSON(c *gin.Context, status int, message string, data interface{}, err interface{}) {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
	c.JSON(status, response)
}
