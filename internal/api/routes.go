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
	mux.Handle("POST /assignments/{id}/complete", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerToggleAssignmentCompletion)))
	mux.Handle("POST /assignments/{id}/reschedule", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRescheduleAssignment)))

	// users
	mux.Handle("POST /users", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerCreateUser)))
	mux.Handle("GET /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetUserByID)))
	mux.Handle("PUT /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerEditUser)))
	mux.Handle("GET /users", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetHouseholdUsers)))
	mux.Handle("DELETE /users/{id}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerDeleteUser)))
	mux.Handle("GET /me", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetMe)))
	mux.Handle("PUT /me/password", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerChangePassword)))
	mux.Handle("PUT /users/{id}/reset-password", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerResetPassword)))

	// households
	mux.Handle("POST /households/{hhid}/users", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerAddUserToHousehold)))
	mux.Handle("GET /households/{hhid}/users", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetHouseholdUsers)))
	mux.Handle("GET /households/{hhid}/users/active", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetActiveHouseholdUsers)))
	mux.Handle("PUT /households/{hhid}/users/{uid}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerEditHouseholdUser)))
	mux.Handle("POST /households/{hhid}/users/{uid}/activate", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerActivateUser)))
	mux.Handle("POST /households/{hhid}/users/{uid}/deactivate", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerDeactivateHouseholdUser)))
	mux.Handle("DELETE /households/{hhid}/users/{uid}", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRemoveUserFromHousehold)))

	// admin
	mux.Handle("POST /scheduler/run", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRunScheduler)))
	mux.Handle("POST /balancer/run", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerRunBalancer)))
	mux.Handle("POST /login", http.HandlerFunc(cfg.handlerLogin))
	mux.Handle("POST /logout", http.HandlerFunc(cfg.handlerLogout))
	mux.Handle("POST /bootstrap", http.HandlerFunc(cfg.handlerBootstrap))

	// events
	mux.Handle("GET /events", cfg.middlewareAuth(http.HandlerFunc(cfg.handlerGetEvents)))
}
