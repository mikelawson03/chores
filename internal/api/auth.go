package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/auth"
)

func (cfg *apiCfg) middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tok, err := auth.GetBearerToken(r.Header)
		if err != nil {
			err = fmt.Errorf("%w: error retrieving token", app.ErrUnauthorized)
			RespondWithError(w, err)
			return
		}

		uid, err := auth.ValidateToken(tok, cfg.App.Config.JWTSigninSecret)
		if err != nil {
			err = fmt.Errorf("%w: error validating token", app.ErrUnauthorized)
			RespondWithError(w, err)
			return
		}

		user, err := cfg.App.Store.GetUserByID(ctx, uid)

		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("%w: user not found", app.ErrUnauthorized)
			RespondWithError(w, err)
			return
		}

		if err != nil {
			RespondWithError(w, err)
			return
		}

		ctx = auth.WithUser(ctx, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
