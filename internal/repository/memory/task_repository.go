package memory

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (repo *Repository) CreateTask(ctx context.Context, listID int, task models.Task) models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]

	return list.AddTask(task)
}

func (repo *Repository) GetTasks(ctx context.Context, listID int) []models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	tasks := make([]models.Task, 0, len(list.Tasks))
	for _, task := range list.Tasks {
		tasks = append(tasks, *task)
	}

	return sortTasksByID(tasks)
}

func (repo *Repository) GetTask(ctx context.Context, listID, taskID int) (models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	task, ok := list.Tasks[taskID]
	if !ok {
		return models.Task{}, errorsx.ErrTaskNotFound
	}

	return *task, nil
}

func (repo *Repository) EditTask(ctx context.Context, listID, taskID int, task models.Task) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	list.Tasks[taskID] = &task

	return nil
}

/*
	func (repo *Repository) CompleteTask(ctx context.Context, listID, taskID int) {
		repo.mu.Lock()
		defer repo.mu.Unlock()

		list := repo.Lists[listID]
		task := list.Tasks[taskID]
		task.IsCompleted = true
	}

	func (repo *Repository) UncompleteTask(ctx context.Context, listID, taskID int) {
		repo.mu.Lock()
		defer repo.mu.Unlock()

		list := repo.Lists[listID]
		task := list.Tasks[taskID]
		task.IsCompleted = false
	}
*/

func (repo *Repository) DeleteTask(ctx context.Context, listID, taskID int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list := repo.Lists[listID]
	delete(list.Tasks, taskID)
}
