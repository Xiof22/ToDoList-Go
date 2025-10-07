package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type ToDoRepository struct {
	Tasks  []models.Task
	mu     sync.Mutex
	nextID int
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{nextID: 1}
}

func (repo *ToDoRepository) Create(title, description string) {
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
}
