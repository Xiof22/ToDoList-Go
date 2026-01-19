package dto

import "github.com/Xiof22/ToDoList/internal/models"

type List struct {
	OwnerID     string `json:"owner_id,omitempty"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	TasksCount  int    `json:"tasks_count"`
}

func ToListDTO(l models.List, withOwnerID bool) List {
	var ownerID string
	if withOwnerID {
		ownerID = l.OwnerID.String()
	}

	return List{
		OwnerID:     ownerID,
		ID:          l.ID.String(),
		Title:       l.Title,
		Description: l.Description,
		TasksCount:  len(l.Tasks),
	}
}

func ToListDTOs(lists []models.List, withOwnerID bool) []List {
	listDTOs := make([]List, len(lists))
	for i, l := range lists {
		listDTOs[i] = ToListDTO(l, withOwnerID)
	}

	return listDTOs
}
