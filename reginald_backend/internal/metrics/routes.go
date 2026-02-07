package metrics

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", MetricsHandler)
	return mux
}
