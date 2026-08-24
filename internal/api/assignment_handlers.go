package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/domain"
)

type AssignmentRequest struct {
	TemplateID     string     `json:"templateId"`
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
	err := d.Decode(target)
	if err != nil {
		return domain.ErrInvalidRequest
	}

	return nil
}

func (cfg *apiCfg) handlerCreateAssignment(w http.ResponseWriter, r *http.Request) {
	req := AssignmentRequest{}
	err := decodeRequest(r, &req)
	if err != nil {
		RespondWithError(w, err)
		return
	}
	ctx := r.Context()

	assignment, err := cfg.App.CreateAssignmentFromTemplate(ctx, app.CreateAsssignmentRequest{
		TemplateID:     req.TemplateID,
		AssignedUserID: req.AssignedUserID,
		Instructions:   req.Instructions,
		DueDate:        req.DueDate,
		ScheduledFor:   req.ScheduledFor,
	})
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
		err := fmt.Errorf("%w: may only specify one assignment filter", domain.ErrInvalidRequest)
		RespondWithError(w, err)
		return
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
		return
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

	ctx := r.Context()
	res, err := cfg.App.EditAssignment(ctx, app.EditAssignmentRequest{
		ID:             assignmentID,
		AssignedUserID: req.AssignedUserID,
		ScheduledFor:   req.ScheduledFor,
		Notes:          req.Notes,
		Completed:      req.Completed,
		Canceled:       req.Canceled,
	})
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}

func (cfg *apiCfg) handlerToggleAssignmentCompletion(w http.ResponseWriter, r *http.Request) {
	assignmentID := r.PathValue("id")
	ctx := r.Context()

	res, err := cfg.App.ToggleAssignmentCompletion(ctx, assignmentID)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}
