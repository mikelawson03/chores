package api

import (
	"net/http"
	"os"
)

func middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-User-ID")
		if user != os.Getenv("USER") {
			RespondWithJSON(w, http.StatusUnauthorized, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
