package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"rate-limiter/controller"
	"rate-limiter/service"
	"rate-limiter/store"
)

func main() {
	// Initialize Store
	memStore := store.NewMemoryStore()

	// Initialize Rate Limiter
	rateLimiter := service.NewRateLimiter(memStore, 5, time.Minute)

	// Initialize Handler
	handler := controller.NewHandler(rateLimiter)

	// setup Gin Router
	r := gin.Default()

	// Routes
	r.POST("/request", handler.HandleRequest)
	r.GET("/stats", handler.GetStats)

	// configure port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
