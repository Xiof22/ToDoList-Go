package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/errorsx"
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

func (repo *Repository) GetTasks(ctx context.Context) ([]models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks := make([]models.Task, len(repo.Tasks))
	i := 0
	for _, t := range repo.Tasks {
		tasks[i] = *t
		i++
	}

	return tasks, nil
}

func (repo *Repository) GetTask(ctx context.Context, taskID models.TaskID) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, ok := repo.Tasks[taskID]
	if !ok {
		return models.Task{}, errorsx.ErrTaskNotFound
	}

	return *task, nil
}
