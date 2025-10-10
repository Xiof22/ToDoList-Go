package errorsx

import "errors"

var (
	ErrTaskNotFound       = errors.New("Task not found")
	ErrInvalidTaskID      = errors.New("Invalid task ID")
	ErrAlreadyCompleted   = errors.New("Task is already completed")
	ErrAlreadyUncompleted = errors.New("Task is already uncompleted")
)
