package main

import (
	"github.com/agrawal-2005/iiita_lab/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to IIITA Lab API"})
	})

	r.Run(":8080")
}