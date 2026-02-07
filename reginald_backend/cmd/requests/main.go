package main

import (
	"log"
	"net/http"
	"reginald_backend/internal/requests"
	"reginald_backend/pkg/database"
)

func main() {
	// Initialize database
	if _, err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	database.InitDB()

	router := requests.SetupRoutes()

	log.Println("Requests service running on :8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
