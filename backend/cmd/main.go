package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/agrawal-2005/iiita_lab/backend/pkg/config"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()

	// Registering routes
	routes.SetupAuthRoutes(r)
	routes.SetupProfileRoutes(r)
	routes.SetupProjectRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3301"
	}
	serverAddr := ":" + port

	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	go func() {
		log.Printf("🚀 Server is running on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🔻 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server shutdown error: %v", err)
	}

	log.Println("✅ Server gracefully stopped.")
}
