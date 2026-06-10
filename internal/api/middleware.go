package api

import (
	"database/sql"
	"errors"
	"net/http"
)

func (cfg *apiCfg) middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get("X-User-ID")

		_, err := cfg.App.GetUserByID(ctx, id)

		if errors.Is(err, sql.ErrNoRows) {
			RespondWithJSON(w, http.StatusUnauthorized, nil)
			return
		}

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Database error: ", err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
