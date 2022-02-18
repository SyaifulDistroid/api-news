package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, code int, status interface{}, message interface{}, result interface{}) {
	c.JSON(code, gin.H{"status": status, "message": message, "result": result})
}
