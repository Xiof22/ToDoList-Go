package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

func (repo *Repository) CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	task := &models.Task{
		ID:          list.NextID,
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		Deadline:    req.Deadline.Value,
	}

	list.Tasks[list.NextID] = task
	list.NextID++
	return *task
}

func (repo *Repository) GetTasks(ctx context.Context, req dto.ListIdentifier) []models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ID]
	tasks := make([]models.Task, 0, len(list.Tasks))
	for _, task := range list.Tasks {
		tasks = append(tasks, *task)
	}

	return sortTasks(tasks)
}

func (repo *Repository) GetTask(ctx context.Context, req dto.TaskIdentifier) (*models.Task, bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	task, found := list.Tasks[req.TaskID]
	return task, found
}

func (repo *Repository) EditTask(ctx context.Context, req dto.EditTaskRequest) models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	task := list.Tasks[req.TaskID]

	task.Title = req.Title
	task.Description = req.Description
	task.Deadline = req.Deadline.Value
	task.UpdatedAt = time.Now()

	return *task
}

func (repo *Repository) CompleteTask(ctx context.Context, req dto.TaskIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	task := list.Tasks[req.TaskID]
	task.IsCompleted = true
}

func (repo *Repository) UncompleteTask(ctx context.Context, req dto.TaskIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	task := list.Tasks[req.TaskID]
	task.IsCompleted = false
}

func (repo *Repository) DeleteTask(ctx context.Context, req dto.TaskIdentifier) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[req.ListID]
	delete(list.Tasks, req.TaskID)
}
