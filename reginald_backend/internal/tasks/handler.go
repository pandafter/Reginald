// HTTP HANDLER FOR TASKS
package tasks

import (
	"encoding/json"
	"net/http"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	case http.MethodPatch, http.MethodPut:
		updateTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("X-User-Email")
	role := r.Header.Get("X-User-Role")

	var tasks []Task
	var err error

	if role == "admin" {
		// Admins see everything
		tasks, err = ListTasks("")
	} else {
		// Users only see their own tasks
		tasks, err = ListTasks(email)
	}

	if err != nil {
		http.Error(w, "Could not fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-User-Role")
	if role != "admin" {
		http.Error(w, "Only admins can create tasks", http.StatusForbidden)
		return
	}

	var req CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = ValidateCreateTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := CreateTask(req)
	if err != nil {
		http.Error(w, "Could not create task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     string `json:"id"`
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	email := r.Header.Get("X-User-Email")
	err = UpdateTaskByID(req.ID, req.Title, req.Status, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		// Try to decode from body if not in query
		var req struct {
			ID string `json:"id"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		id = req.ID
	}

	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	email := r.Header.Get("X-User-Email")
	err := DeleteTaskByID(id, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	email := r.Header.Get("X-User-Email")
	err = ChangeTaskStatus(req.ID, req.Status, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Status updated successfully"})
}
