package service

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateList(title, description string) error {
	return svc.repo.CreateList(title, description)
}

func (svc *Service) GetLists() ([]models.List, error) {
	return svc.repo.GetLists()
}

func (svc *Service) GetList(listID int) (*models.List, error) {
	return svc.repo.GetList(listID)
}

func (svc *Service) EditList(listID int, title, description string) error {
	return svc.repo.EditList(listID, title, description)
}

func (svc *Service) DeleteList(listID int) error {
	return svc.repo.DeleteList(listID)
}
