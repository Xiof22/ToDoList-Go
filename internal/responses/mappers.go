package responses

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"net/http"
)

func MapError(err error) int {
	switch {
	case errors.Is(err, errorsx.ErrInvalidListID),
		errors.Is(err, errorsx.ErrInvalidTaskID),
		errors.Is(err, errorsx.ErrAlreadyCompleted),
		errors.Is(err, errorsx.ErrAlreadyUncompleted),
		errors.Is(err, errorsx.ErrDeadlineBeforeCreation):
		return http.StatusBadRequest

	case errors.Is(err, errorsx.ErrTaskNotFound),
		errors.Is(err, errorsx.ErrListNotFound):
		return http.StatusNotFound

	case errors.Is(err, errorsx.ErrHashPassword):
		return http.StatusInternalServerError

	case errors.Is(err, errorsx.ErrInvalidCredentials),
		errors.Is(err, errorsx.ErrUserNotFound):
		return http.StatusUnauthorized

	case errors.Is(err, errorsx.ErrEmailRegistered):
		return http.StatusConflict

	case errors.Is(err, errorsx.ErrForbidden):
		return http.StatusForbidden

	default:
		return http.StatusInternalServerError
	}
}
