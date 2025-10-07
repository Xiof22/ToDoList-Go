package memory

import (
	"context"
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

func (repo *Repository) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Tasks[task.ID] = &task

	return task, nil
}
