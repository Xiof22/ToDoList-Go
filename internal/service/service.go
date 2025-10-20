package service

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
	"time"
)

type ToDoService struct {
	repo *repository.ToDoRepository
}

func NewToDoService(repo *repository.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (svc *ToDoService) CreateTask(title, description string, deadline time.Time) error {
	return svc.repo.Create(title, description, deadline)
}

func (svc *ToDoService) GetTasks() []models.Task {
	return svc.repo.GetAll()
}

func (svc *ToDoService) GetTask(id int) *models.Task {
	return svc.repo.Get(id)
}

func (svc *ToDoService) EditTask(id int, title, description string, deadline time.Time) error {
	return svc.repo.Edit(id, title, description, deadline)
}

func (svc *ToDoService) CompleteTask(id int) error {
	return svc.repo.Complete(id)
}

func (svc *ToDoService) UncompleteTask(id int) error {
	return svc.repo.Uncomplete(id)
}

func (svc *ToDoService) DeleteTask(id int) error {
	return svc.repo.Delete(id)
}
