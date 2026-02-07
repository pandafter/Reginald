package users

import (
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/register", CreateUserHandler)
	mux.HandleFunc("/users/login", LoginHandler)
	mux.HandleFunc("/users/", GetUsersHandler)
	mux.HandleFunc("/users", GetUsersHandler)

	return mux
}
