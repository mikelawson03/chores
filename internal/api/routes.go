package api

import (
	"net/http"
)

func (cfg *apiCfg) RegisterRoutes(mux *http.ServeMux) {
	// chore templates
	mux.Handle("GET /chore-templates", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetChoreTemplates)))
	mux.Handle("GET /chore-templates/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetChoreTemplateByID)))
	mux.Handle("POST /chore-templates", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerAddChoreTemplate)))
	mux.Handle("PUT /chore-templates/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerEditChoreTemplate)))
	mux.Handle("DELETE /chore-templates/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerDeleteChoreTemplate)))

	// assigned chores
	mux.Handle("POST /assignments", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerCreateAssignment)))
	mux.Handle("GET /assignments", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetAssignments)))
	mux.Handle("GET /assignments/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetAssignmentByID)))
	mux.Handle("PUT /assignments/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerEditAssignment)))

	// users
	mux.Handle("POST /users", http.HandlerFunc(cfg.handlerAddUser))
	mux.Handle("GET /users", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetUsers)))
	mux.Handle("GET /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetUserByID)))
	mux.Handle("PUT /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerEditUser)))
	mux.Handle("DELETE /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerDeleteUser)))

	// admin
	mux.Handle("POST /scheduler/run", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRunScheduler)))
	mux.Handle("POST /balancer/run", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRunBalancer)))
	mux.Handle("POST /login", http.HandlerFunc(cfg.handlerLogin))
}
