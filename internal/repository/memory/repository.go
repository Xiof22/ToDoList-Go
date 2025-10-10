package memory

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

var (
	ErrNotFound           = errors.New("Task not found")
	ErrAlreadyCompleted   = errors.New("Task is already completed")
	ErrAlreadyUncompleted = errors.New("Task is already uncompleted")
)

type Repository struct {
	Tasks  []models.Task
	mu     sync.Mutex
	nextID int
}

func New() *Repository {
	return &Repository{nextID: 1}
}

func (repo *Repository) Create(title, description string) error {
	task := models.Task{
		ID:          repo.nextID,
		Title:       title,
		Description: description,
		IsCompleted: false,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Tasks = append(repo.Tasks, task)
	repo.nextID++

	return nil
}

func (repo *Repository) GetAll() ([]models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.Tasks, nil
}

func (repo *Repository) Get(id int) (*models.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for index, task := range repo.Tasks {
		if task.ID == id {
			return &repo.Tasks[index], nil
		}
	}

	return nil, nil
}

func (repo *Repository) Edit(id int, title, description string) error {
	task, _ := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.Title = title
	task.Description = description

	return nil
}

func (repo *Repository) Complete(id int) error {
	task, _ := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	if task.IsCompleted {
		return ErrAlreadyCompleted
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.IsCompleted = true
	return nil
}

func (repo *Repository) Uncomplete(id int) error {
	task, _ := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	if !task.IsCompleted {
		return ErrAlreadyUncompleted
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.IsCompleted = false
	return nil
}
