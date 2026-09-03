package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
)

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

type EditHouseholdUserRequest struct {
	Role        string `json:"role"`
	DisplayName string `json:"displayName"`
	ColorOption int    `json:"colorOption"`
	IsActive    bool   `json:"isActive"`
	HouseholdId string `json:"householdId"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ResetPasswordRequest struct {
	TempPassword string `json:"tempPassword"`
}

func CreateUserRequest(r *http.Request) (*UserRequest, error) {
	d := json.NewDecoder(r.Body)
	req := &UserRequest{}

	err := d.Decode(req)
	if err != nil {
		return &UserRequest{}, domain.ErrInvalidRequest
	}

	return req, nil
}

func (cfg *apiCfg) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userReq, err := CreateUserRequest(r)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	user, err := cfg.App.CreateNewUser(ctx, userReq.Username, userReq.FirstName, userReq.Password)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, user)
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

func (cfg *apiCfg) handlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	user, err := cfg.App.GetUserByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		RespondWithError(w, err)
		return
	}

	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerEditUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	updateUser, err := CreateUserRequest(r)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	user, err := cfg.App.EditUser(ctx, id, updateUser.Username, updateUser.FirstName)

	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
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
		UserId:      userId,
		HouseholdId: householdId,
	})

	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)

}

func (cfg *apiCfg) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	err := cfg.App.DeleteUser(ctx, id)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *apiCfg) handlerLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	d := json.NewDecoder(r.Body)
	req := &LoginRequest{}

	err := d.Decode(req)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	result, err := cfg.App.LoginUser(ctx, req.Username, req.Password)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	resp := LoginResponse{
		Token: result.Token,
	}

	RespondWithJSON(w, http.StatusOK, resp)

}

func (cfg *apiCfg) handlerGetMe(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticatedUser(r.Context())
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	d := json.NewDecoder(r.Body)
	pwReq := &ChangePasswordRequest{}

	err := d.Decode(pwReq)
	if err != nil {
		RespondWithError(w, domain.ErrInvalidRequest)
		return
	}

	err = cfg.App.ChangePassword(ctx, pwReq.OldPassword, pwReq.NewPassword)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)

}

func (cfg *apiCfg) handlerResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	d := json.NewDecoder(r.Body)
	req := &ResetPasswordRequest{}

	err := d.Decode(req)
	if err != nil {
		RespondWithError(w, domain.ErrInvalidRequest)
		return
	}

	err = cfg.App.ResetPassword(ctx, id, req.TempPassword)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)
}
