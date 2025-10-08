package memory

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

type Repository struct {
	Tasks  []models.Task
	mu     sync.Mutex
	nextID int
}

func New() *Repository {
	return &Repository{nextID: 1}
}

func (repo *Repository) Create(title, description string) {
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
