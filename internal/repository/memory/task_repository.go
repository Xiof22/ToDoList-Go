package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

func (repo *Repository) CreateTask(listID int, title, description string, deadline time.Time) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	now := time.Now()
	if deadline.Before(now) && !deadline.IsZero() {
		return ErrInvalidDeadline
	}

	task := &models.Task{
		ID:          list.NextID,
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   now,
		Deadline:    deadline,
	}

	list.Tasks[list.NextID] = task
	list.NextID++

	return nil
}

func (repo *Repository) GetTasks(listID int) ([]models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return nil, ErrListNotFound
	}

	taskCount := len(list.Tasks)
	tasks := make([]models.Task, 0, taskCount)

	for _, task := range list.Tasks {
		tasks = append(tasks, *task)
	}

	return sortTasks(tasks), nil
}

func (repo *Repository) GetTask(listID, taskID int) (*models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return nil, ErrListNotFound
	}

	return list.Tasks[taskID], nil
}

func (repo *Repository) EditTask(listID, taskID int, title, description string, deadline time.Time) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	task, found := list.Tasks[taskID]
	if !found {
		return ErrTaskNotFound
	}

	if deadline.Before(task.CreatedAt) && !deadline.IsZero() {
		return ErrInvalidDeadline
	}

	task.Title = title
	task.Description = description
	task.Deadline = deadline
	task.UpdatedAt = time.Now()

	return nil
}

func (repo *Repository) CompleteTask(listID, taskID int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	task, found := list.Tasks[taskID]
	if !found {
		return ErrTaskNotFound
	}

	if task.IsCompleted {
		return ErrAlreadyCompleted
	}

	task.IsCompleted = true

	return nil
}

func (repo *Repository) UncompleteTask(listID, taskID int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	task, found := list.Tasks[taskID]
	if !found {
		return ErrTaskNotFound
	}

	if !task.IsCompleted {
		return ErrAlreadyUncompleted
	}

	task.IsCompleted = false

	return nil
}

func (repo *Repository) DeleteTask(listID, taskID int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	list, found := repo.Lists[listID]
	if !found {
		return ErrListNotFound
	}

	_, found = list.Tasks[taskID]
	if !found {
		return ErrTaskNotFound
	}

	delete(list.Tasks, taskID)
	return nil
}
