package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/agrawal-2005/iiita_lab/pkg/config"
)

func main() {
	config.Connect()
	
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})		
	})

	port := ":3301"
	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		log.Println("🚀 Server is running on port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🔻 Shutting down server...")
	if err := server.Close(); err != nil {
		log.Fatalf("❌ Server shutdown error: %v", err)
	}

	log.Println("✅ Server gracefully stopped.")
}