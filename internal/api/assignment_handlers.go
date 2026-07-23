package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
)

type AssignmentRequest struct {
	ChoreID        string     `json:"choreId"`
	AssignedUserID string     `json:"assignedUserId"`
	Instructions   string     `json:"instructions"`
	DueDate        *time.Time `json:"dueDate"`
	ScheduledFor   *time.Time `json:"scheduledFor"`
}

type EditAssignmentRequest struct {
	AssignedUserID string     `json:"userId"`
	ScheduledFor   *time.Time `json:"scheduledFor"`
	Notes          string     `json:"notes"`
	Completed      bool       `json:"completed"`
	Canceled       bool       `json:"canceled"`
}

func decodeRequest(r *http.Request, target any) error {
	d := json.NewDecoder(r.Body)

	return d.Decode(target)
}

func (cfg *apiCfg) handlerCreateAssignmentForUser(w http.ResponseWriter, r *http.Request) {
	// Decode request into assignment request struct (make helper for this)
	req := AssignmentRequest{}
	err := decodeRequest(r, &req)
	if err != nil {
		RespondWithError(w, err)
		return
	}
	// call app layer for authorization and assignment creation

	requesterID := r.Header.Get("X-User-ID")
	ctx := r.Context()

	assignment, err := cfg.App.CreateAssignmentForUser(ctx, requesterID, req.ChoreID, req.AssignedUserID, req.Instructions, req.DueDate, req.ScheduledFor)

	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, assignment)
}

func (cfg *apiCfg) handlerGetAssignments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	template_id := r.URL.Query().Get("template_id")
	user_id := r.URL.Query().Get("user_id")

	if template_id != "" && user_id != "" {
		err := errors.New("may only specify one assignment filter")
		RespondWithError(w, err)
	}

	var (
		assignments []domain.Assignment
		err         error
	)

	switch {
	case template_id != "":
		assignments, err = cfg.App.GetAssignmentsByTemplateID(ctx, template_id)
	case user_id != "":
		assignments, err = cfg.App.GetAssignmentsByUserID(ctx, user_id)
	default:
		assignments, err = cfg.App.GetAllAssignments(ctx)
	}

	if err != nil {
		RespondWithError(w, err)
	}

	RespondWithJSON(w, http.StatusOK, assignments)
}

func (cfg *apiCfg) handlerGetAssignmentByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	assignment, err := cfg.App.GetAssignmentByID(ctx, id)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}

func (cfg *apiCfg) handlerEditAssignment(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	req := EditAssignmentRequest{}
	err := decodeRequest(r, &req)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	requesterID := r.Header.Get("X-User-ID")
	ctx := r.Context()
	res, err := cfg.App.EditAssignment(ctx, requesterID, assignmentID, req.AssignedUserID, req.Notes, req.Canceled, req.Completed, req.ScheduledFor)
	if err != nil {
		RespondWithError(w, err)
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
		RespondWithError(w, err)
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
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignment)
}
