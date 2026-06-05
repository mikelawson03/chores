package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errorResponse{
		Error: fmt.Sprint(msg, err),
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.WriteHeader(code)
	if code != http.StatusNoContent {
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			log.Printf("Error marshaling response: %s\n", err)
			w.WriteHeader(500)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
}
