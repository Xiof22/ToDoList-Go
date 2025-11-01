package service

import "errors"

var (
	ErrTaskNotFound           = errors.New("Task not found")
	ErrAlreadyCompleted       = errors.New("Task is already completed")
	ErrAlreadyUncompleted     = errors.New("Task is already uncompleted")
	ErrDeadlineBeforeCreation = errors.New("Deadline must be after task creation time")
	ErrListNotFound           = errors.New("List not found")
)
