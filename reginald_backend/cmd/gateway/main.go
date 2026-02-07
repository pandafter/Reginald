package main

import (
	"log"
	"net/http"

	"reginald_backend/internal/auth"
	"reginald_backend/internal/metrics"
	"reginald_backend/pkg/database"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize database for metrics
	if _, err := database.Connect(); err != nil {
		log.Printf("Warning: Monitoring metrics will not be available: %v", err)
	}

	mux := http.NewServeMux()

	// API Routes prefixed with /api to avoid collision with frontend routes
	// Public API
	mux.HandleFunc("/api/health", HealthCheckHandler)
	mux.HandleFunc("/api/auth/login", ProxyToAuth)
	mux.HandleFunc("/api/auth/logout", ProxyToAuth)
	mux.HandleFunc("/api/users/login", ProxyToUsers)
	mux.HandleFunc("/api/users/register", ProxyToUsers)

	// Protected API
	mux.Handle("/api/auth/verify", auth.JWTMiddleware(http.HandlerFunc(ProxyToAuth)))
	mux.Handle("/api/users/", auth.JWTMiddleware(http.HandlerFunc(ProxyToUsers)))
	mux.Handle("/api/requests/", auth.JWTMiddleware(http.HandlerFunc(ProxyToRequests)))
	mux.Handle("/api/tasks/", auth.JWTMiddleware(http.HandlerFunc(ProxyToTasks)))
	mux.Handle("/api/monitoring/", auth.JWTMiddleware(http.HandlerFunc(metrics.MetricsHandler)))
	mux.Handle("/api/metrics/latency", auth.JWTMiddleware(http.HandlerFunc(metrics.LatencyHandler)))

	// Strip /api prefix before proxying or handling if needed
	// Actually, the proxy functions in proxy.go use r.URL.Path directly.
	// We need to either update proxy.go or use http.StripPrefix here.

	apiHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a catch-all for /api/
		mux.ServeHTTP(w, r)
	})

	// Wrap everything with CORS and Security
	handler := corsMiddleware(apiHandler)
	handler = SecurityMiddleware(handler)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
