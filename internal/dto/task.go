package dto

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	IsCompleted bool       `json:"completed"`
	Deadline    *time.Time `json:"deadline,omitempty"`
}

func ToTaskDTO(t *models.Task) *Task {
	if t == nil {
		return nil
	}

	var deadline *time.Time = nil
	if !t.Deadline.IsZero() {
		deadline = &t.Deadline
	}

	return &Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsCompleted: t.IsCompleted,
		Deadline:    deadline,
	}
}

func ToTaskDTOs(tasks []models.Task) []Task {
	taskDTOs := make([]Task, len(tasks))
	for i, t := range tasks {
		taskDTOs[i] = *ToTaskDTO(&t)
	}

	return taskDTOs
}
