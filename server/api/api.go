package api

import (
	v1 "sample-project/server/api/v1"

	"github.com/gin-gonic/gin"
)

// SetupAPIRoutes function
func SetupAPIRoutes(router *gin.RouterGroup) {
	v1.SetupV1Routes(router.Group("/v1"))
}
