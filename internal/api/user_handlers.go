package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

type UserRequest struct {
	Username  string `json:"username"`
	Role      string `json:"role"`
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

type LoginResponse struct {
	User  UserResponse `json:"id"`
	Token string       `json:"token"`
}

func CreateUserRequest(r *http.Request) (*UserRequest, error) {
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
	userReq, err := CreateUserRequest(r)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	user, err := cfg.App.CreateNewUser(ctx, userReq.Username, userReq.Role, userReq.FirstName)
	if err != nil {
		RespondWithError(w, err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, user)
}

func (cfg *apiCfg) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := cfg.App.GetAllUsers(ctx)
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

	user, err := cfg.App.EditUser(ctx, id, updateUser.Username, updateUser.Role, updateUser.FirstName)

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

	if errors.Is(err, sql.ErrNoRows) {
		RespondWithError(w, err)
		return
	}

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
		User: UserResponse{
			ID:        result.User.ID,
			Username:  result.User.Username,
			FirstName: result.User.FirstName,
			Role:      string(result.User.Role),
		},
		Token: result.Token,
	}

	RespondWithJSON(w, http.StatusOK, resp)

}
