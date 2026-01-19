package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type Repository struct {
	mu         sync.Mutex
	Lists      map[int]*models.List
	Users      map[int]*models.User
	listNextID int
	userNextID int
}

func New() *Repository {
	return &Repository{
		Lists:      make(map[int]*models.List),
		Users:      make(map[int]*models.User),
		listNextID: 1,
		userNextID: 1,
	}
}
