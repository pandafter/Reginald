package main

import (
	"log"
	"net/http"

	"reginald_backend/internal/users"
	"reginald_backend/pkg/database"
)

func main() {
	// Initialize database
	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	database.InitDB()

	// Bootstrap admin user
	users.BootstrapAdmin()

	router := users.SetupRoutes()

	log.Println("Users service running on :8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
