package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"reginald_backend/pkg/database"

	"github.com/google/uuid"
)

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isSuspicious(r) {
			logSecurityAlert(r)
			http.Error(w, "Security Alert: Malicious Request Detected", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isSuspicious(r *http.Request) bool {
	// 1. Check URL Query
	decodedQuery, _ := url.QueryUnescape(r.URL.RawQuery)
	if checkSQLInjection(decodedQuery) {
		return true
	}

	// 2. Check Path
	decodedPath, _ := url.QueryUnescape(r.URL.Path)
	if checkSQLInjection(decodedPath) {
		return true
	}

	// 3. Check Body (Post Data)
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				// Important: Restore body for the next handlers/proxy
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				bodyStr := string(bodyBytes)
				if checkSQLInjection(bodyStr) {
					return true
				}
			}
		}
	}

	return false
}

func checkSQLInjection(s string) bool {
	if s == "" {
		return false
	}
	// Expanded patterns to catch common SQLi including the user's test case
	// Patterns:
	// - ' OR '1'='1
	// - 1' or '1' = '1
	// - ; DROP TABLE
	// - --
	// - UNION SELECT
	patterns := []string{
		`(?i)('|\d)\s*or\s*('|\d)?\s*1\s*'?\s*=\s*'?\s*1`, // Matches 1 or 1=1, ' or '1'='1, etc.
		`(?i);\s*drop\s+table`,
		`--`,
		`(?i)union\s+select`,
		`(?i)('|\d)\s*or\s*true`,
	}

	for _, p := range patterns {
		re := regexp.MustCompile(p)
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

func logSecurityAlert(r *http.Request) {
	id := uuid.New().String()[:16]
	userEmail := r.Header.Get("X-User-Email")
	if userEmail == "" {
		userEmail = "unknown/attacker"
	}

	ip := getIP(r)

	query := `INSERT INTO audit_logs (id, user_email, action, resource, resource_id, status_code, duration_ms, timestamp) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := database.DB.Exec(query, id, userEmail, "SECURITY_ALERT", r.URL.String(), ip, 403, 0, time.Now())
	if err != nil {
		log.Printf("Failed to log security alert: %v", err)
	}
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	} else {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ip = strings.Split(ip, ",")[0]
	}
	// Basic cleanup for IPv6 local address
	if ip == "[" || ip == "::1" {
		return "127.0.0.1"
	}
	return ip
}
