package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type Repository struct {
	mu    sync.Mutex
	Tasks map[models.TaskID]*models.Task
}

func New() *Repository {
	return &Repository{
		Tasks: make(map[models.TaskID]*models.Task),
	}
}
