package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mikelawson03/chores/internal/domain"
)

func RespondWithError(w http.ResponseWriter, err error) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	switch {
	case errors.Is(err, domain.ErrUnauthorized):
		RespondWithJSON(w, http.StatusUnauthorized, errorResponse{
			Error: "Unauthorized",
		})

	case errors.Is(err, domain.ErrForbidden):
		RespondWithJSON(w, http.StatusForbidden, errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrInvalidCredentials):
		RespondWithJSON(w, http.StatusUnauthorized, errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrNotFound):
		RespondWithJSON(w, http.StatusNotFound, errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrInvalidRole),
		errors.Is(err, domain.ErrInvalidCadence),
		errors.Is(err, domain.ErrPasswordRequired),
		errors.Is(err, domain.ErrPasswordTooShort),
		errors.Is(err, domain.ErrInvalidColorOption),
		errors.Is(err, domain.ErrInvalidDisplayName),
		errors.Is(err, domain.ErrInvalidRequest):
		RespondWithJSON(w, http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrConflict),
		errors.Is(err, domain.ErrHouseholdFull):
		RespondWithJSON(w, http.StatusConflict, errorResponse{
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
