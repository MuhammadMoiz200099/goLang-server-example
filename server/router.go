package main

import (
	"sample-project/server/api"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func InitializeRoutes(router *gin.Engine) {
	api.SetupAPIRoutes(router.Group("/api"))
}
