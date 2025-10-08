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

func (repo *ToDoRepository) GetAll() []models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.Tasks
}

func (repo *ToDoRepository) Get(id int) *models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for index, task := range repo.Tasks {
		if task.ID == id {
			return &repo.Tasks[index]
		}
	}

	return nil
}
