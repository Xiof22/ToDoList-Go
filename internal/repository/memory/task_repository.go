package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

func (repo *Repository) CreateTask(ctx context.Context, listID models.ListID, task models.Task) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]

	task.CreatedAt = time.Now()

	list.Tasks[task.ID] = &task

	return task, nil
}

func (repo *Repository) GetTasks(ctx context.Context, listID models.ListID) ([]models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	tasks := make([]models.Task, 0, len(list.Tasks))
	for _, task := range list.Tasks {
		tasks = append(tasks, *task)
	}

	return sortTasksByCreationTime(tasks), nil
}

func (repo *Repository) GetTask(ctx context.Context, listID models.ListID, taskID models.TaskID) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	task, ok := list.Tasks[taskID]
	if !ok {
		return models.Task{}, errorsx.ErrTaskNotFound
	}

	return *task, nil
}

func (repo *Repository) EditTask(ctx context.Context, listID models.ListID, taskID models.TaskID, task models.Task) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	now := time.Now()
	task.UpdatedAt = &now
	list.Tasks[taskID] = &task

	return task, nil
}

func (repo *Repository) DeleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	delete(list.Tasks, taskID)

	return nil
}
