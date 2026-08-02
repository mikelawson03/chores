package api

import (
	"encoding/json"
	"net/http"

	"github.com/mikelawson03/chores/internal/domain"
)

type ChoreTemplateRequest struct {
	Name         string `json:"name"`
	Cadence      string `json:"cadence"`
	Assignee     string `json:"assignee"`
	Instructions string `json:"instructions"`
	Duration     *int   `json:"duration"`
}

func CreateChoreTeplateRequest(r *http.Request) (*ChoreTemplateRequest, error) {
	d := json.NewDecoder(r.Body)
	req := &ChoreTemplateRequest{}

	err := d.Decode(req)
	if err != nil {
		return &ChoreTemplateRequest{}, err
	}

	return req, nil
}

func (cfg *apiCfg) handlerGetChoreTemplates(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chores, err := cfg.App.GetChoreTemplates(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}
	if len(chores) == 0 {
		RespondWithJSON(w, http.StatusOK, []domain.ChoreTemplate{})
		return
	}
	RespondWithJSON(w, http.StatusOK, chores)
}

func (cfg *apiCfg) handlerAddChoreTemplate(w http.ResponseWriter, r *http.Request) {
	req, err := CreateChoreTeplateRequest(r)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	ctx := r.Context()
	chore, err := cfg.App.CreateChoreTemplate(ctx, req.Name, req.Cadence, req.Assignee, req.Instructions, req.Duration)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, chore)
}

func (cfg *apiCfg) handlerGetChoreTemplateByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	c, err := cfg.App.GetChoreTemplateByID(ctx, id)
	if err != nil {
		RespondWithError(w, err)
		return
	}
	RespondWithJSON(w, http.StatusOK, c)
}

func (cfg *apiCfg) handlerEditChoreTemplate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	req, err := CreateChoreTeplateRequest(r)
	if err != nil {
		RespondWithError(w, err)
		return
	}
	ctx := r.Context()
	c, err := cfg.App.EditChoreTemplate(ctx, id, req.Name, req.Cadence, req.Assignee, req.Instructions, req.Duration)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, c)
}

func (cfg *apiCfg) handlerDeleteChoreTemplate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()

	err := cfg.App.DeleteChoreTemplate(ctx, id)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
