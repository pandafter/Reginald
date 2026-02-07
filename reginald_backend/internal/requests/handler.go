package requests

import (
	"encoding/json"
	"net/http"
)

func RequestsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getRequests(w, r)
	case http.MethodPost:
		createRequest(w, r)
	case http.MethodPatch, http.MethodPut:
		processRequest(w, r)
	case http.MethodDelete:
		deleteRequest(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case "HEALTH": // Simple health check
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getRequests(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-User-Email")
	role := r.Header.Get("X-User-Role")

	list, err := ListRequests(email, role)
	if err != nil {
		http.Error(w, "Could not fetch requests", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func createRequest(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-User-Email")
	var dto CreateRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateCreateRequest(dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto.CreatedBy = email

	req, err := CreateRequest(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-User-Role")
	adminEmail := r.Header.Get("X-User-Email")

	if role != "admin" {
		http.Error(w, "Only admins can process requests", http.StatusForbidden)
		return
	}

	var body struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err := ProcessRequest(body.ID, body.Status, adminEmail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Request updated"})
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-User-Role")
	if role != "admin" {
		http.Error(w, "Only admins can delete requests", http.StatusForbidden)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Request ID is required", http.StatusBadRequest)
		return
	}

	err := DeleteRequest(id)
	if err != nil {
		http.Error(w, "Could not delete request", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
