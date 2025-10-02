package repository

import (
	"sync"
	"github.com/Xiof22/ToDoList/internal/models"
)

type ToDoRepository struct {
	mu sync.Mutex
	Tasks []models.Task
	nextID int
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{ nextID : 1 }
}

func (repo *ToDoRepository) CreateTask(title, description string) {
	task := models.Task{
		ID : repo.nextID,
		Title : title,
		Description : description,
		IsCompleted : false,
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.Tasks = append(repo.Tasks, task)
	repo.nextID++
}
