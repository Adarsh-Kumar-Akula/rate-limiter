package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"rate-limiter/controller"
	"rate-limiter/service"
	"rate-limiter/store"
)

func main() {
	// --- Initialize Store ---
	memStore := store.NewMemoryStore()

	// --- Initialize Rate Limiter ---
	rateLimiter := service.NewRateLimiter(memStore, 5, time.Minute)

	// --- Initialize Handler ---
	handler := controller.NewHandler(rateLimiter)

	// --- Setup Gin Router ---
	r := gin.Default()

	// --- Routes ---
	r.POST("/request", handler.HandleRequest)
	r.GET("/stats", handler.GetStats)

	// --- Start Server ---
	port := ":8080"
	log.Printf("Server running on %s\n", port)

	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
