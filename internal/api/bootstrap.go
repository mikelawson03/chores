package api

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiCfg) handlerBootstrap(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	d := json.NewDecoder(r.Body)
	userReq := &UserRequest{}

	err := d.Decode(userReq)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	result, err := cfg.App.Bootstrap(ctx, userReq.Username, userReq.FirstName, userReq.Password)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, result)

}
