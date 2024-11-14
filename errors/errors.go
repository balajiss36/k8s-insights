package errors

import (
	"github.com/gin-gonic/gin"
)

// NewError creates a new error
func NewError(c *gin.Context, err error, statusCode int, details interface{}) {
	c.JSON(statusCode, gin.H{
		"error":   err.Error(),
		"message": details,
	})
}
