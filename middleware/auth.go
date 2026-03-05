package middleware

import (
	"net/http"

	"github.com/AaravShetty15/go-todo-app/config"
)

func BasicAuth(next http.Handler) http.Handler {

	cfg := config.LoadConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()

		if !ok || username != cfg.AuthUser || password != cfg.AuthPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}