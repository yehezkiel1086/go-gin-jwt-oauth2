package util

import "github.com/gin-gonic/gin"

func ResponseHandler(c *gin.Context, status int, isErr bool, msg string) {
	if isErr {
		c.AbortWithStatusJSON(status, gin.H{
			"error": msg,
		})
		// c.JSON(status, gin.H{
		// 	"error": msg,
		// })
	} else {
		c.JSON(status, gin.H{
			"message": msg,
		})
	}
}