package errorsx

import (
	"errors"
	"fmt"
)

var (
	ErrWriteJSON     = errors.New("Failed to write JSON")
	ErrMissingJSON   = errors.New("Missing JSON")
	ErrInvalidTaskID = errors.New("Invalid task ID")
	ErrTaskNotFound  = errors.New("Task not found")
)

func ErrValidation(field, rule string) error {
	return fmt.Errorf("Field '%s' doesn't match the rule '%s'", field, rule)
}
