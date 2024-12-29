package main

import (
	"log"
	"net/http"

	"github.com/arunsingh28/api/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.RegisterRoutes(router)

	port := "8080"

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
