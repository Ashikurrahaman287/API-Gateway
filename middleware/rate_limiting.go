package middleware

import (
	"github.com/julienschmidt/httprouter" // Import httprouter for rate limiting
	"net/http"
)

// RateLimitingMiddleware is a middleware function for rate limiting
func RateLimitingMiddleware(next http.Handler) http.Handler {
	limiter := httprouter.NewRateLimiter(1, 5) // Allow 1 request per second, burst of 5

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Limit() {
			// Request rate limit exceeded
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
