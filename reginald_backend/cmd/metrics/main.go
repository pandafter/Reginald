package main

import (
	"log"
	"net/http"

	"reginald_backend/internal/metrics"
)

func main() {
	router := metrics.SetupRoutes()

	log.Println("Metrics service running on :8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
