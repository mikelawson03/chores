package api

import "net/http"

type apiCfg struct{}

func NewApiConfig() apiCfg {
	return apiCfg{}
}

func (cfg *apiCfg) RegisterRoutes(mux *http.ServeMux) {
	// chore templates
	mux.Handle("GET /chore-templates", middlewareAuth(http.HandlerFunc(cfg.getChores)))
}
