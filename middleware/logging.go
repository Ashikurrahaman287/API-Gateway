package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function for logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log request information
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}
