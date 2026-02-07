package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func getServiceURL(envVar, fallback string) string {
	val := os.Getenv(envVar)
	if val == "" {
		return fallback
	}
	return val
}

func ProxyToTasks(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	url := getServiceURL("TASKS_SERVICE_URL", "http://localhost:8082") + path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}
	proxyRequest(w, r, url)
}

func ProxyToUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	url := getServiceURL("USERS_SERVICE_URL", "http://localhost:8083") + path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}
	proxyRequest(w, r, url)
}

func ProxyToAuth(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	url := getServiceURL("AUTH_SERVICE_URL", "http://localhost:8081") + path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}
	proxyRequest(w, r, url)
}

func ProxyToRequests(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")
	url := getServiceURL("REQUESTS_SERVICE_URL", "http://localhost:8084") + path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}
	proxyRequest(w, r, url)
}

func proxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// Copy headers
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
