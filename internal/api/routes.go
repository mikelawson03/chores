package api

import (
	"net/http"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/domain"
)

type apiCfg struct {
	App *app.App
}

func NewApiConfig() apiCfg {
	return apiCfg{
		App: &app.App{
			Templates:   make(map[string]domain.ChoreTemplate),
			Assignments: make([]domain.Assignment, 0),
		},
	}
}

func (cfg *apiCfg) RegisterRoutes(mux *http.ServeMux) {
	// chore templates
	mux.Handle("GET /chore-templates", middlewareAuth(http.HandlerFunc(cfg.handlerGetChoreTemplates)))
	mux.Handle("GET /chore-templates/{id}", middlewareAuth(http.HandlerFunc(cfg.handlerGetChoreTemplateByID)))
	mux.Handle("POST /chore-templates", middlewareAuth(http.HandlerFunc(cfg.handlerAddChoreTemplate)))
	mux.Handle("PUT /chore-templates/{id}", middlewareAuth(http.HandlerFunc(cfg.handlerEditChoreTemplate)))
	mux.Handle("DELETE /chore-templates/{id}", middlewareAuth(http.HandlerFunc(cfg.handlerDeleteChoreTemplate)))

	// assigned chores
	mux.Handle("POST /assignments", middlewareAuth(http.HandlerFunc(cfg.handlerCreateAssignmentForUser)))
}
