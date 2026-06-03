package api

import "net/http"

func handlerTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (cfg *apiCfg) getChores(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
