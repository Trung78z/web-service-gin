package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Response(c *gin.Context, status int, success bool, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
	})
}
