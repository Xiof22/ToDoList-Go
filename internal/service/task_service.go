package service

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

func (svc *Service) CreateTask(listID int, title, description string, deadline time.Time) error {
	return svc.repo.CreateTask(listID, title, description, deadline)
}

func (svc *Service) GetTasks(listID int) ([]models.Task, error) {
	return svc.repo.GetTasks(listID)
}

func (svc *Service) GetTask(listID, taskID int) (*models.Task, error) {
	return svc.repo.GetTask(listID, taskID)
}

func (svc *Service) EditTask(listID, taskID int, title, description string, deadline time.Time) error {
	return svc.repo.EditTask(listID, taskID, title, description, deadline)
}

func (svc *Service) CompleteTask(listID, taskID int) error {
	return svc.repo.CompleteTask(listID, taskID)
}

func (svc *Service) UncompleteTask(listID, taskID int) error {
	return svc.repo.UncompleteTask(listID, taskID)
}

func (svc *Service) DeleteTask(listID, taskID int) error {
	return svc.repo.DeleteTask(listID, taskID)
}
