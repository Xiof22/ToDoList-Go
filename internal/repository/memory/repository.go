package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type Repository struct {
	mu     sync.Mutex
	Lists  map[int]*models.List
	nextID int
}

func New() *Repository {
	return &Repository{
		Lists:  make(map[int]*models.List),
		nextID: 1,
	}
}
