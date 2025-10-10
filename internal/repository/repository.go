package repository

import (
	"sync"
	"errors"
	"github.com/Xiof22/ToDoList/internal/models"
)

var ErrNotFound = errors.New("Task not found")

type ToDoRepository struct {
	mu sync.Mutex
	Tasks []models.Task
	nextID int
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{ nextID : 1 }
}

func (repo *ToDoRepository) Create(title, description string) {
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

func (repo *ToDoRepository) Edit(id int, title, description string) error {
	task := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.Title = title
	task.Description = description
	return nil
}

func (repo *ToDoRepository) Complete(id int) error {
	task := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	if task.IsCompleted {
		return errors.New("Task is already completed")
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.IsCompleted = true
	return nil
}

func (repo *ToDoRepository) Uncomplete(id int) error {
	task := repo.Get(id)
	if task == nil {
		return ErrNotFound
	}

	if !task.IsCompleted {
		return errors.New("Task is already uncompleted")
	}

	repo.mu.Lock()
	defer repo.mu.Unlock()

	task.IsCompleted = false
	return nil
}

func (repo *ToDoRepository) Delete(id int) error {
	for index, task := range repo.Tasks {
		if task.ID == id {
			repo.mu.Lock()
			defer repo.mu.Unlock()

			repo.Tasks = append(repo.Tasks[:index], repo.Tasks[index + 1:]...)
			return nil
		}
	}

	return ErrNotFound
}
