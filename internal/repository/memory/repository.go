package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

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

func (repo *Repository) CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task {
	task := &models.Task{
		ID:          repo.nextID,
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: false,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Tasks[repo.nextID] = task
	repo.nextID++

	return *task
}
