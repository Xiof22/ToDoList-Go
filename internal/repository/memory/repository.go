package memory

import "github.com/Xiof22/ToDoList/internal/models"

type Repository struct {
	mu     sync.Mutex
	Tasks  map[int]*models.Task
	nextID int
}

func New() *Repository {
	return &Repository{
		Tasks:  make(map[int]*models.Task),
		nextID: 1,
	}
}
