package api

import (
	"net/http"
)

func (cfg *apiCfg) RegisterRoutes(mux *http.ServeMux) {
	// chore templates
	mux.Handle("GET /chore-templates", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetChoreTemplates)))
	mux.Handle("GET /chore-templates/{id}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetChoreTemplateByID)))
	mux.Handle("POST /chore-templates", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerAddChoreTemplate)))
	mux.Handle("PUT /chore-templates/{id}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerEditChoreTemplate)))
	mux.Handle("DELETE /chore-templates/{id}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerDeleteChoreTemplate)))

	// assigned chores
	mux.Handle("POST /assignments", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerCreateAssignment)))
	mux.Handle("GET /assignments", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetAssignments)))
	mux.Handle("GET /assignments/{id}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetAssignmentByID)))
	mux.Handle("PUT /assignments/{id}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerEditAssignment)))
	mux.Handle("POST /assignments/{id}/complete", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerToggleAssignmentCompletion)))
	mux.Handle("POST /assignments/{id}/reschedule", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerRescheduleAssignment)))

	// users
	mux.Handle("POST /users", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerCreateUser)))
	mux.Handle("GET /users/{id}", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerGetUserByID)))
	mux.Handle("PUT /users/{id}", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerEditUser)))
	mux.Handle("GET /users", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerGetHouseholdUsers)))
	mux.Handle("DELETE /users/{id}", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerDeleteUser)))
	mux.Handle("GET /me", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetMe)))
	mux.Handle("PUT /me/password", cfg.middlewareUserAuth(http.HandlerFunc(cfg.handlerChangePassword)))
	mux.Handle("PUT /users/{id}/reset-password", cfg.middlewareHouseholdAuth(http.HandlerFunc(cfg.handlerResetPassword)))

	// households
	mux.Handle("POST /households/{hhid}/users", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerAddUserToHousehold)))
	mux.Handle("GET /households/{hhid}/users", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetHouseholdUsers)))
	mux.Handle("GET /households/{hhid}/users/active", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetActiveHouseholdUsers)))
	mux.Handle("PUT /households/{hhid}/users/{uid}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerEditHouseholdUser)))
	mux.Handle("POST /households/{hhid}/users/{uid}/activate", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerActivateUser)))
	mux.Handle("POST /households/{hhid}/users/{uid}/deactivate", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerDeactivateHouseholdUser)))
	mux.Handle("DELETE /households/{hhid}/users/{uid}", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerRemoveUserFromHousehold)))

	// admin
	mux.Handle("POST /scheduler/run", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerRunScheduler)))
	mux.Handle("POST /balancer/run", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerRunBalancer)))
	mux.Handle("POST /login", http.HandlerFunc(cfg.handlerLogin))
	mux.Handle("POST /logout", http.HandlerFunc(cfg.handlerLogout))
	mux.Handle("POST /bootstrap", http.HandlerFunc(cfg.handlerBootstrap))

	// events
	mux.Handle("GET /events", cfg.requireHouseholdAuth(http.HandlerFunc(cfg.handlerGetEvents)))
}
