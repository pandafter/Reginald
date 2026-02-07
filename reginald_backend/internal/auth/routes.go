package auth

import "net/http"

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", LoginHandler)
	mux.HandleFunc("/auth/logout", LogoutHandler)
	mux.HandleFunc("/auth/verify", VerifyHandler)
	return mux
}
