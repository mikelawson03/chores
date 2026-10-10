package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
)

func (cfg *apiCfg) requireHouseholdAuth(next http.Handler) http.Handler {
	return cfg.middlewareUserAuth(
		cfg.middlewareHouseholdAuth(next),
	)
}

func (cfg *apiCfg) middlewareUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var tok string

		cookie, err := r.Cookie("auth")

		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				err = fmt.Errorf("%w: auth cookie not found", domain.ErrUnauthorized)
				RespondWithError(w, err)
				return
			}

			RespondWithError(w, err)
			return
		}

		tok = cookie.Value

		uid, err := auth.ValidateToken(tok, cfg.App.Config.JWTSigninSecret)
		if err != nil {
			err = fmt.Errorf("%w: error validating token", domain.ErrUnauthorized)
			RespondWithError(w, err)
			return
		}

		ctx = auth.WithUserID(ctx, uid)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (cfg *apiCfg) middlewareHouseholdAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid, ok := auth.UserIDFromContext(ctx)
		if !ok {
			err := fmt.Errorf("%w: authenticated user ID missing from context", domain.ErrUnauthorized)
			log.Println(err)
			RespondWithError(w, err)
			return
		}

		user, err := cfg.App.Store.GetHouseholdUserByID(ctx, uid, domain.DefaultHouseholdID)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				err = fmt.Errorf("%w: user not found", domain.ErrUnauthorized)
				RespondWithError(w, err)
				return
			}

			RespondWithError(w, err)
			return
		}

		if !user.IsActive {
			err = fmt.Errorf("%w: household membership is inactive", domain.ErrInactiveHouseholdUser)
			RespondWithError(w, err)
			return
		}

		ctx = auth.WithHouseholdUser(ctx, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
