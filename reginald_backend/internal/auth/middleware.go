package auth

import (
	"context"
	"net/http"
	"strings"

	"reginald_backend/pkg/utils"
)

type contextKey string

const UserEmailKey contextKey = "user_email"
const UserRoleKey contextKey = "user_role"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := ""

		// 1. Try Authorization Header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}

		// 2. Try Cookie if header is missing
		if token == "" {
			cookie, err := r.Cookie("token")
			if err == nil {
				token = cookie.Value
			}
		}

		if token == "" {
			http.Error(w, "unauthorized: token missing", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, "unauthorized: invalid or expired token", http.StatusUnauthorized)
			return
		}

		id, _ := claims["id"].(string)
		name, _ := claims["name"].(string)
		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)

		ctx := context.WithValue(r.Context(), UserEmailKey, email)
		ctx = context.WithValue(ctx, UserRoleKey, role)

		// Propagate user info to downstream services via headers
		r.Header.Set("X-User-ID", id)
		r.Header.Set("X-User-Name", name)
		r.Header.Set("X-User-Email", email)
		r.Header.Set("X-User-Role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
