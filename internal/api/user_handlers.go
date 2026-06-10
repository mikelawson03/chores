package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username"`
}

func CreateNewUserRequest(r *http.Request) (*UserRequest, error) {
	d := json.NewDecoder(r.Body)
	req := &UserRequest{}

	err := d.Decode(req)
	if err != nil {
		return &UserRequest{}, err
	}

	return req, nil
}

func (cfg *apiCfg) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userReq, err := CreateNewUserRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error", err)
		return
	}

	user, err := cfg.App.CreateNewUser(ctx, userReq.Username)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error creating new user", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, user)
}

func (cfg *apiCfg) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := cfg.App.GetAllUsers(ctx)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error retrieving users", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiCfg) handlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	user, err := cfg.App.GetUserByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		RespondWithError(w, http.StatusNotFound, "User not found", nil)
		return
	}

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Database error", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerEditUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")
	updatedUser, err := CreateNewUserRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Request Error: ", err)
		return
	}

	user, err := cfg.App.EditUser(ctx, id, updatedUser.Username, requesterID)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error updating user: ", err)
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiCfg) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	requesterID := r.Header.Get("X-User-ID")
	err := cfg.App.DeleteUser(ctx, id, requesterID)

	if errors.Is(err, sql.ErrNoRows) {
		RespondWithError(w, http.StatusNotFound, "User not found", nil)
		return
	}

	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Error deleting user: ", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
