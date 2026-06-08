package api

import (
	"encoding/json"
	"net/http"

	"github.com/mikelawson03/chores/internal/domain"
)

type ChoreTemplateRequest struct {
	Name     string `json:"name"`
	Cadence  string `json:"cadence"`
	Shared   *bool  `json:"shared"`
	Assignee string `json:"assignee"`
	Duration *int   `json:"duration"`
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
	// TODO: replace with logic to call list of chores once they have been created
	chores := cfg.App.GetChoreTemplates()
	if len(chores) == 0 {
		RespondWithJSON(w, http.StatusOK, []domain.ChoreTemplate{})
		return
	}
	RespondWithJSON(w, http.StatusOK, chores)
}

func (cfg *apiCfg) handlerAddChoreTemplate(w http.ResponseWriter, r *http.Request) {
	req, err := CreateChoreTeplateRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error: ", err)
		return
	}

	chore, err := cfg.App.CreateChoreTemplate(req.Name, req.Cadence, req.Shared, req.Assignee, req.Duration)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error adding chore: ", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, chore)
}

func (cfg *apiCfg) handlerGetChoreTemplateByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, err := cfg.App.GetChoreByID(id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "", err)
		return
	}
	RespondWithJSON(w, http.StatusOK, c)
}

func (cfg *apiCfg) handlerEditChoreTemplate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	req, err := CreateChoreTeplateRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error decoding JSON: ", err)
		return
	}

	c, err := cfg.App.EditChoreTemplate(id, req.Name, req.Cadence, req.Shared, req.Assignee, req.Duration)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Error updating chore: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, c)
}

func (cfg *apiCfg) handlerDeleteChoreTemplate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := cfg.App.DeleteChoreTemplate(id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Error deleting chore: ", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *apiCfg) handlerGetAssignmentsByTemplateID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	assignments, err := cfg.App.GetAssignmentsByTemplateID(id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Error retrieving assignments: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, assignments)
}
