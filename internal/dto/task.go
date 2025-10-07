package dto

import "github.com/Xiof22/ToDoList/internal/models"

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	IsCompleted bool   `json:"completed"`
}

func ToTaskDTO(t models.Task) Task {
	return Task{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		IsCompleted: t.IsCompleted,
	}
}
