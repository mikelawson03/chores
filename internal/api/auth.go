package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/mikelawson03/chores/internal/auth"
)

func (cfg *apiCfg) middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// id := r.Header.Get("X-User-ID")

		tok, err := auth.GetBearerToken(r.Header)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Error retrieving token: ", err)
			return
		}

		uid, err := auth.ValidateToken(tok, cfg.App.Config.JWTSigninSecret)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Error validating token: ", err)
			return
		}

		user, err := cfg.App.GetUserByID(ctx, uid)

		if errors.Is(err, sql.ErrNoRows) {
			RespondWithJSON(w, http.StatusUnauthorized, nil)
			return
		}

		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Database error: ", err)
			return
		}

		ctx = auth.WithUser(ctx, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
