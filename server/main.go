package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	InitializeRoutes(router)

	router.Run(":" + port)
}
