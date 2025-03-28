package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/q-summitLabs/cefi/tree/main/backend/internal/handlers"
	"github.com/q-summitLabs/cefi/tree/main/backend/internal/middleware"
)

func main() {

	// Init Gin Router
	router := gin.Default()

	// Apply CORS middlewaree from internal/middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	router.GET("/api/data", handlers.GetData)

	// Start the server on port 8080
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}

}
