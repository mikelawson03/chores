package domain

import "errors"

var (
	ErrConflict              = errors.New("conflict")
	ErrForbidden             = errors.New("forbidden")
	ErrHouseholdFull         = errors.New("household_full")
	ErrInactiveHouseholdUser = errors.New("inactive_household_user")
	ErrInvalidCadence        = errors.New("invalid_cadence")
	ErrInvalidColorOption    = errors.New("invalid_color_option")
	ErrInvalidCredentials    = errors.New("invalid_credentials")
	ErrInvalidDisplayName    = errors.New("invalid_display_name")
	ErrInvalidRequest        = errors.New("invalid_request")
	ErrInvalidRole           = errors.New("invalid_role")
	ErrNotFound              = errors.New("not_found")
	ErrPasswordRequired      = errors.New("password_required")
	ErrPasswordTooShort      = errors.New("password_too_short")
	ErrUnauthorized          = errors.New("unauthorized")
)
