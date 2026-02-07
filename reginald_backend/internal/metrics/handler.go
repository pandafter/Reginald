package metrics

import (
	"encoding/json"
	"log"
	"net/http"
	"reginald_backend/pkg/database"

	"github.com/google/uuid"
)

type LatencyReport struct {
	Endpoint   string `json:"endpoint"`
	Method     string `json:"method"`
	DurationMs int    `json:"duration_ms"`
	StatusCode int    `json:"status_code"`
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	role := r.Header.Get("X-User-Role")
	if !IsAdmin(role) {
		http.Error(w, ErrUnauthorized.Error(), http.StatusForbidden)
		return
	}

	metrics, err := GetSystemMetrics()
	if err != nil {
		http.Error(w, "Could not fetch metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func LatencyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var report LatencyReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userEmail := r.Header.Get("X-User-Email")
	if userEmail == "" {
		userEmail = "anonymous"
	}

	// Log to audit_logs table
	query := `INSERT INTO audit_logs (id, user_email, action, resource, resource_id, status_code, duration_ms) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	id := uuid.New().String()[:16]
	_, err = database.DB.Exec(query, id, userEmail, report.Method, report.Endpoint, "", report.StatusCode, report.DurationMs)
	if err != nil {
		log.Printf("Error logging latency: %v", err)
		// Don't fail the request, just log
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "recorded"})
}
