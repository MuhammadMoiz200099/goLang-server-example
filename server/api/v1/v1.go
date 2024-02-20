package v1

import (
	"github.com/gin-gonic/gin"
)

// SetupV1Routes function
func SetupV1Routes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "V1 API endpoint",
		})
	})
}
