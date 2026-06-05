package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type AssignmentRequest struct {
	ChoreID        string     `json:"chore_id"`
	AssignedUserID string     `json:"assigned_user_id"`
	ScheduleDate   *time.Time `json:"schedule_date"`
}

func CreateNewAssignmentRequest(r *http.Request) (*AssignmentRequest, error) {
	d := json.NewDecoder(r.Body)
	req := &AssignmentRequest{}

	err := d.Decode(req)
	if err != nil {
		return &AssignmentRequest{}, err
	}

	return req, nil
}

func (cfg *apiCfg) handlerCreateAssignmentForUser(w http.ResponseWriter, r *http.Request) {
	// Decode request into assignment request struct (make helper for this)
	req, err := CreateNewAssignmentRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error: ", err)
		return
	}
	// call app layer for authorization and assignment creation

	requesterID := r.Header.Get("X-User-ID")

	assignment, err := cfg.App.CreateAssignmentForUser(requesterID, req.ChoreID, req.AssignedUserID, req.ScheduleDate)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error assigning chore: ", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, assignment)
}
