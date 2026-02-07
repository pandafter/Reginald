package main

import (
	"log"
	"net/http"

	"reginald_backend/internal/auth"
	"reginald_backend/pkg/database"
)

func main() {
	// Initialize database
	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	database.InitDB()

	router := auth.SetupRoutes()

	log.Println("Auth service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
