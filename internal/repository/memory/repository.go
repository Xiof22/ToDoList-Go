package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type Repository struct {
	mu    sync.Mutex
	Lists map[models.ListID]*models.List
	Users map[models.UserID]*models.User
}

func New() *Repository {
	return &Repository{
		Lists: make(map[models.ListID]*models.List),
		Users: make(map[models.UserID]*models.User),
	}
}
