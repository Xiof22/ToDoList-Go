package dto

import "github.com/Xiof22/ToDoList/internal/models"

type List struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	TasksCount  int    `json:"tasks_count"`
}

func ToListDTO(l *models.List) *List {
	if l == nil {
		return nil
	}

	return &List{
		ID:          l.ID,
		Title:       l.Title,
		Description: l.Description,
		TasksCount:  len(l.Tasks),
	}
}

func ToListDTOs(lists []models.List) []List {
	listDTOs := make([]List, len(lists))
	for i, t := range lists {
		listDTOs[i] = *ToListDTO(&t)
	}

	return listDTOs
}
