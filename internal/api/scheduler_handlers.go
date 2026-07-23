package api

import "net/http"

func (cfg *apiCfg) handlerRunScheduler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := cfg.App.RunScheduler(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, "scheduler completed successfully")
}
