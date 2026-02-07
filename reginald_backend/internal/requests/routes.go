package requests

import (
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/requests", RequestsHandler)
	mux.HandleFunc("/requests/", RequestsHandler)

	return mux
}
