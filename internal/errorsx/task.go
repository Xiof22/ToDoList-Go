package errorsx

import "errors"

var (
	ErrTaskNotFound           = errors.New("Task not found")
	ErrInvalidTaskID          = errors.New("Invalid task ID")
	ErrAlreadyCompleted       = errors.New("Task is already completed")
	ErrAlreadyUncompleted     = errors.New("Task is already uncompleted")
	ErrUnmarshalDeadline      = errors.New("Deadline unmarshalling error")
	ErrInvalidDeadlineFormat  = errors.New("Unexpected deadline format")
	ErrDeadlineBeforeCreation = errors.New("Deadline must be after task creation time")
)
