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

type EditAssignmentRequest struct {
	AssignedUserID string     `json:"assigned_user_id"`
	ScheduleDate   *time.Time `json:"schedule_date"`
}

func decodeRequest(r *http.Request, target any) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	return d.Decode(target)
}

func (cfg *apiCfg) handlerCreateAssignmentForUser(w http.ResponseWriter, r *http.Request) {
	// Decode request into assignment request struct (make helper for this)
	req := AssignmentRequest{}
	err := decodeRequest(r, &req)
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
	ctx := r.Context()
	assignments, err := cfg.App.GetAllAssignments(ctx)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error retrieving assignments:", err)
	}

	RespondWithJSON(w, http.StatusOK, assignments)
}

func (cfg *apiCfg) handlerGetAssignmentByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	assignment, err := cfg.App.GetAssignmentByID(ctx, id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error retrieving assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}

func (cfg *apiCfg) handlerEditAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	req := EditAssignmentRequest{}
	err := decodeRequest(r, &req)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error: ", err)
		return
	}

	requesterID := r.Header.Get("X-User-ID")
	ctx := r.Context()
	res, err := cfg.App.EditAssignment(ctx, requesterID, assignmentID, req.AssignedUserID, req.ScheduleDate)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error editing assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}

func (cfg *apiCfg) handlerCancelAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")
	ctx := r.Context()

	assignment, err := cfg.App.CancelAssignment(ctx, assignmentID, requesterID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error deleting assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}

func (cfg *apiCfg) handlerCompleteAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")
	ctx := r.Context()

	assignment, err := cfg.App.CompleteAssignment(ctx, assignmentID, requesterID)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error completing assignment: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}
