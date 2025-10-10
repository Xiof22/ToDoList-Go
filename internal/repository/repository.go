package repository

import (
	"errors"
	"github.com/Xiof22/ToDoList/internal/models"
	"sync"
)

var (
	ErrNotFound = errors.New("Task not found")
)

type ToDoRepository struct {
	mu     sync.Mutex
	Tasks  map[int]models.Task
	nextID int
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		Tasks:  make(map[int]models.Task),
		nextID: 1,
	}
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

	repo.Tasks[repo.nextID] = task
	repo.nextID++
}

func (repo *ToDoRepository) GetAll() []models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	taskCount := len(repo.Tasks)
	tasks := make([]models.Task, 0, taskCount)

	for _, task := range repo.Tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (repo *ToDoRepository) Get(id int) *models.Task {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return nil
	}

	return &task
}

func (repo *ToDoRepository) Edit(id int, title, description string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	task.Title = title
	task.Description = description
	repo.Tasks[id] = task

	return nil
}

func (repo *ToDoRepository) Complete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	if task.IsCompleted {
		return errors.New("Task is already completed")
	}

	task.IsCompleted = true
	repo.Tasks[id] = task

	return nil
}

func (repo *ToDoRepository) Uncomplete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	task, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	if !task.IsCompleted {
		return errors.New("Task is already uncompleted")
	}

	task.IsCompleted = false
	repo.Tasks[id] = task

	return nil
}

func (repo *ToDoRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, found := repo.Tasks[id]
	if !found {
		return ErrNotFound
	}

	delete(repo.Tasks, id)
	return nil
}
