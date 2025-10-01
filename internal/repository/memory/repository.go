package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository struct {
	Tasks  []models.Task
	nextID int
}

func New() *Repository {
	return &Repository{nextID: 1}
}
