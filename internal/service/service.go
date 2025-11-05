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

func (svc *Service) CreateTask(title, description string) error {
	return svc.repo.Create(title, description)
}

func (svc *Service) GetTasks() ([]models.Task, error) {
	return svc.repo.GetAll()
}

func (svc *Service) GetTask(id int) (*models.Task, error) {
	return svc.repo.Get(id)
}

func (svc *Service) EditTask(id int, title, description string) error {
	return svc.repo.Edit(id, title, description)
}

func (svc *Service) CompleteTask(id int) error {
	return svc.repo.Complete(id)
}

func (svc *Service) UncompleteTask(id int) error {
	return svc.repo.Uncomplete(id)
}

func (svc *Service) DeleteTask(id int) error {
	return svc.repo.Delete(id)
}
