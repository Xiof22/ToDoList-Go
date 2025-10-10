package responses

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"net/http"
)

func MapError(err error) int {
	switch {

	case errors.Is(err, errorsx.ErrTaskNotFound):
		return http.StatusNotFound

	case errors.Is(err, errorsx.ErrAlreadyCompleted),
		errors.Is(err, errorsx.ErrAlreadyUncompleted):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError

	}
}
