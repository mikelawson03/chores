package api

import (
	"encoding/json"
	"net/http"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/domain"
)

type AddHouseholdUserRequest struct {
	Role        string `json:"role"`
	DisplayName string `json:"displayName"`
	ColorOption int    `json:"colorOption"`
	IsActive    bool   `json:"isActive"`
	UserID      string `json:"userID"`
	HouseholdID string `json:"householdID"`
}

type EditHouseholdUserRequest struct {
	Role        string `json:"role"`
	DisplayName string `json:"displayName"`
	ColorOption int    `json:"colorOption"`
	IsActive    bool   `json:"isActive"`
}

func (cfg *apiCfg) handlerAddUserToHousehold(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hhID := r.PathValue("hhid")

	d := json.NewDecoder(r.Body)
	req := &AddHouseholdUserRequest{}

	err := d.Decode(req)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	hhUser, err := cfg.App.AddUserToHousehold(ctx, app.HouseholdUserRequest{
		Role:        req.Role,
		DisplayName: req.DisplayName,
		ColorOption: req.ColorOption,
		IsActive:    req.IsActive,
		UserID:      req.UserID,
		HouseholdID: hhID,
	})
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, hhUser)

}

func (cfg *apiCfg) handlerGetHouseholdUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := cfg.App.GetHouseholdUsers(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiCfg) handlerGetActiveHouseholdUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := cfg.App.GetActiveHouseholdUsers(ctx)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiCfg) handlerEditHouseholdUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	householdId := r.PathValue("hhid")
	userId := r.PathValue("uid")

	d := json.NewDecoder(r.Body)
	req := &EditHouseholdUserRequest{}

	err := d.Decode(req)
	if err != nil {
		RespondWithError(w, domain.ErrInvalidRequest)
		return
	}

	user, err := cfg.App.EditHouseholdUser(ctx, app.HouseholdUserRequest{
		Role:        req.Role,
		DisplayName: req.DisplayName,
		ColorOption: req.ColorOption,
		IsActive:    req.IsActive,
		UserID:      userId,
		HouseholdID: householdId,
	})

	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)

}

func (cfg *apiCfg) handlerActivateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	householdId := r.PathValue("hhid")
	userId := r.PathValue("uid")

	user, err := cfg.App.ActivateHouseholdUser(ctx, householdId, userId)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerDeactivateHouseholdUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	householdId := r.PathValue("hhid")
	userId := r.PathValue("uid")

	user, err := cfg.App.DeactivateHouseholdUser(ctx, householdId, userId)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerRemoveUserFromHousehold(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	householdId := r.PathValue("hhid")
	userId := r.PathValue("uid")

	err := cfg.App.RemoveUserFromHousehold(ctx, householdId, userId)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
