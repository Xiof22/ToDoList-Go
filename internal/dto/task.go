package dto

import "github.com/Xiof22/ToDoList/internal/models"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	IsCompleted bool   `json:"completed"`
}

func ToTaskDTO(t *models.Task) *Task {
	if t == nil {
		return nil
	}

	return &Task{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsCompleted: t.IsCompleted,
	}
}
