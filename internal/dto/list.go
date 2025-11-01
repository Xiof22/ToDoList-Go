package dto

import "github.com/Xiof22/ToDoList/internal/models"

type List struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	TasksCount  int    `json:"tasks_count"`
}

func ToListDTO(l models.List) List {
	return List{
		ID:          l.ID.String(),
		Title:       l.Title,
		Description: l.Description,
		TasksCount:  len(l.Tasks),
	}
}

func ToListDTOs(lists []models.List) []List {
	listDTOs := make([]List, len(lists))
	for i, l := range lists {
		listDTOs[i] = ToListDTO(l)
	}

	return listDTOs
}
