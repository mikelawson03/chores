package domain

import "errors"

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrNotFound           = errors.New("not_found")
	ErrInvalidRole        = errors.New("invalid_role")
	ErrInvalidCadence     = errors.New("invalid_cadence")
	ErrPasswordRequired   = errors.New("password_required")
	ErrPasswordTooShort   = errors.New("password_too_short")
	ErrInvalidRequest     = errors.New("invalid_request")
	ErrConflict           = errors.New("conflict")
)
