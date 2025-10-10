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

func (repo *Repository) GetTasks(ctx context.Context) []models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks := make([]models.Task, len(repo.Tasks))
	i := 0
	for _, t := range repo.Tasks {
		tasks[i] = *t
		i++
	}

	return sortTasks(tasks)
}

func (repo *Repository) GetTask(ctx context.Context, req dto.TaskIdentifier) (*models.Task, bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[req.ID]
	return task, found
}

func (repo *Repository) EditTask(ctx context.Context, req dto.EditTaskRequest) models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task := repo.Tasks[req.ID]
	task.Title = req.Title
	task.Description = req.Description

	return *task
}

func (repo *Repository) CompleteTask(ctx context.Context, req dto.TaskIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task := repo.Tasks[req.ID]
	task.IsCompleted = true
}

func (repo *Repository) UncompleteTask(ctx context.Context, req dto.TaskIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task := repo.Tasks[req.ID]
	task.IsCompleted = false
}
