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
	ctx := r.Context()

	assignment, err := cfg.App.CreateAssignmentForUser(ctx, requesterID, req.ChoreID, req.AssignedUserID, req.ScheduleDate)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error assigning chore: ", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, assignment)
}

func (cfg *apiCfg) handlerGetAllAssignments(w http.ResponseWriter, r *http.Request) {
	assignments := cfg.App.GetAllAssignments()

	RespondWithJSON(w, http.StatusOK, assignments)
}

func (cfg *apiCfg) handlerGetAssignmentByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	assignment, err := cfg.App.GetAssignmentByID(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error retrieving assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}

func (cfg *apiCfg) handlerEditAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	req, err := CreateNewAssignmentRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error: ", err)
		return
	}

	requesterID := r.Header.Get("X-User-ID")
	res, err := cfg.App.EditAssignment(requesterID, assignmentID, req.ChoreID, req.AssignedUserID, req.ScheduleDate)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error editing assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}

func (cfg *apiCfg) handlerDeleteAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")

	err := cfg.App.DeleteAssignment(assignmentID, requesterID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error deleting assignment: ", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *apiCfg) handlerCompleteAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")

	assignment, err := cfg.App.CompleteAssignment(assignmentID, requesterID)

	if err != nil {
		RespondWithError(w, http.StatusInsufficientStorage, "Error completing assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}
