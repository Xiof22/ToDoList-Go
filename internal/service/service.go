package service

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateTask(title, description string) {
	svc.repo.Create(title, description)
}

func (svc *Service) GetTasks() ([]models.Task, error) {
	return svc.repo.GetAll()
}

func (svc *Service) GetTask(id int) (*models.Task, error) {
	return svc.repo.Get(id)
}
