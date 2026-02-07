// Endpoints
package tasks

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks/", TasksHandler)
	mux.HandleFunc("/tasks", TasksHandler)
	mux.HandleFunc("/tasks/status", UpdateStatusHandler)

	return mux
}
