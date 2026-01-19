package errorsx

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSession          = errors.New("Invalid session")
	ErrInvalidListID           = errors.New("Invalid list ID")
	ErrInvalidTaskID           = errors.New("Invalid task ID")
	ErrInvalidUserID           = errors.New("Invalid user ID")
	ErrListNotFound            = errors.New("List not found")
	ErrTaskNotFound            = errors.New("Task not found")
	ErrUserNotFound            = errors.New("User not found")
	ErrAlreadyCompleted        = errors.New("Task is already completed")
	ErrAlreadyUncompleted      = errors.New("Task is already uncompleted")
	ErrDeadlineBeforeCreation  = errors.New("Deadline must be after task creation time")
	ErrForbidden               = errors.New("Content is forbidden")
	ErrUnauthorized            = errors.New("You're not authorized")
	ErrAlreadyAuthorized       = errors.New("You're already authorized")
	ErrEmailRegistered         = errors.New("This email is already registered")
	ErrHashPassword            = errors.New("Failed to hash password")
	ErrSaveSession             = errors.New("Failed to save session")
	ErrInvalidCredentials      = errors.New("Invalid credentials")
	ErrSignToken               = errors.New("Failed to sign token")
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
	ErrInvalidToken            = errors.New("Invalid token")
	ErrMissingToken            = errors.New("Missing token")
	ErrInvalidAuthHeader       = errors.New("Invalid authorization header")
)

func ErrParseURL(key string) error {
	return fmt.Errorf("Failed to parse '%s' from URL", key)
}

func ErrParseContext(key string) error {
	return fmt.Errorf("Failed to parse '%s' from context", key)
}
