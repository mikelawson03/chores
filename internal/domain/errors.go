package domain

import "errors"

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNotFound           = errors.New("not found")
	ErrValidation         = errors.New("invalid request")
	ErrInvalidRole        = errors.New("invalid role")
	ErrInvalidCadence     = errors.New("invalid cadence")
)
