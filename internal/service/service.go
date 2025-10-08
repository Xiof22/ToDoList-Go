package service

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
)

type ToDoService struct {
	repo *repository.ToDoRepository
}

func NewToDoService(repo *repository.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (svc *ToDoService) CreateTask(title, description string) {
	svc.repo.Create(title, description)
}

func (svc *ToDoService) GetTasks() []models.Task {
	return svc.repo.GetAll()
}

func (svc *ToDoService) GetTask(id int) *models.Task {
	return svc.repo.Get(id)
}
