package middleware

import (
	"context"
	"net/http"
)

type AuthMiddleware struct{}

func (a AuthMiddleware) Enable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ctx := context.WithValue(r.Context(), "userid", r.Header.Get("userid"))
		r = r.WithContext(context.WithValue(r.Context(), "token", "123456789"))
		next.ServeHTTP(w, r)
	})
}
