package middleware

import "net/http"

// ContentMiddleware struct
type ContentMiddleware struct{}

// Enable content type
func (c ContentMiddleware) Enable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
