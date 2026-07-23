package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mikelawson03/chores/internal/app"
)

func RespondWithError(w http.ResponseWriter, err error) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	switch {
	case errors.Is(err, app.ErrForbidden):
		RespondWithJSON(w, http.StatusForbidden, errorResponse{
			Error: "Forbidden",
		})
	case errors.Is(err, app.ErrUnauthorized):
		RespondWithJSON(w, http.StatusUnauthorized, errorResponse{
			Error: "Unauthorized",
		})
	case errors.Is(err, app.ErrNotFound):
		RespondWithJSON(w, http.StatusNotFound, errorResponse{
			Error: err.Error(),
		})

	default:
		log.Printf("Unexpected error: %v", err)

		RespondWithJSON(w, http.StatusInternalServerError, errorResponse{
			Error: "Internal Server error",
		})

	}
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
