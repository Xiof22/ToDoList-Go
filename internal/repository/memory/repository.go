package memory

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
	"time"
)

var (
	ErrNotFound           = errors.New("Task not found")
	ErrAlreadyCompleted   = errors.New("Task is already completed")
	ErrAlreadyUncompleted = errors.New("Task is already uncompleted")
	ErrInvalidDeadline    = errors.New("Invalid deadline")
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

func (repo *Repository) Create(title, description string, deadline time.Time) error {
	now := time.Now()

	if deadline.Before(now) && !deadline.IsZero() {
		return ErrInvalidDeadline
	}

	task := &models.Task{
		ID:          repo.nextID,
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   now,
		Deadline:    deadline,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Tasks[repo.nextID] = task
	repo.nextID++

	return nil
}

func (repo *Repository) GetAll() ([]models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	taskCount := len(repo.Tasks)
	tasks := make([]models.Task, 0, taskCount)

	for _, task := range repo.Tasks {
		tasks = append(tasks, *task)
	}

	return sort(tasks), nil
}

func (repo *Repository) Get(id int) (*models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.Tasks[id], nil
}

func (repo *Repository) Edit(id int, title, description string, deadline time.Time) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
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

func (repo *Repository) Complete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	if task.IsCompleted {
		return ErrAlreadyCompleted
	}

	task.IsCompleted = true
	return nil
}

func (repo *Repository) Uncomplete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	if !task.IsCompleted {
		return ErrAlreadyUncompleted
	}

	task.IsCompleted = false

	return nil
}

func (repo *Repository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	delete(repo.Tasks, id)
	return nil
}

func sort(tasks []models.Task) []models.Task {
	for {
		swapped := false

		for i := 0; i < len(tasks)-1; i++ {
			current := tasks[i]
			next := tasks[i+1]

			if current.ID > next.ID {
				tasks[i], tasks[i+1] = next, current
				swapped = true
			}
		}

		if !swapped {
			return tasks
		}
	}
}
