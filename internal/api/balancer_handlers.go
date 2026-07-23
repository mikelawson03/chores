package api

import "net/http"

func (cfg *apiCfg) handlerRunBalancer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := cfg.App.RunBalancer(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, "balancer run successful")
}
