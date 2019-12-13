package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewesteves/tapagguapi/model"
)

// AuthMiddleware type
type AuthMiddleware struct {
	Conn *sql.DB
}

type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

// WithUser set the authenticated user
func WithUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

// GetUser get the authenticated user
func GetUser(ctx context.Context) *model.User {
	user, ok := ctx.Value(userCtxKey).(*model.User)
	if !ok {
		return nil
	}
	return user
}

// Enable authentication by token
func (a AuthMiddleware) Enable(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if guard(r) {
			var user model.User
			var token string
			headerAuth := strings.Split(r.Header.Get("Authorization"), "Bearer")
			if len(headerAuth) > 1 {
				token = headerAuth[1]
			}
			err := a.Conn.QueryRow("SELECT id, name, email, token FROM users WHERE token = $1", strings.TrimSpace(token)).Scan(&user.ID, &user.Name, &user.Email, &user.Token)
			if err != nil {
				http.Error(w, "Please provide the autorization token", http.StatusUnauthorized)
				return
			}
			r = r.WithContext(WithUser(r.Context(), &user))
		}
		next.ServeHTTP(w, r)
	})
}

func guard(r *http.Request) bool {
	accessable := []string{
		"/users-POST",
		"/users/login-POST",
		"/users/recover-POST",
		"/users/confirmation-GET",
	}
	return !contains(accessable, fmt.Sprintf("%s-%s", r.RequestURI, r.Method))
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
