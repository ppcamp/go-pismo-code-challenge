package helpers

import "github.com/gin-gonic/gin"

// GinError returns a json with the msg and its status code.
//
// For status code definition, check http package.
func GinError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
		"status":  status,
	})
}
