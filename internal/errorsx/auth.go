package errorsx

import "errors"

var (
	ErrEmailRegistered    = errors.New("This email is already registered")
	ErrHashPassword       = errors.New("Failed to hash password")
	ErrInvalidCredentials = errors.New("Invalid credentials")
	ErrInvalidSession     = errors.New("Invalid session")
	ErrSaveSession        = errors.New("Failed to save session")
	ErrForbidden          = errors.New("Content is forbidden")
	ErrUnauthorized       = errors.New("You're not authorized")
	ErrAlreadyAuthorized  = errors.New("You're already authorized")
)
