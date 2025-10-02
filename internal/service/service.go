package service

import (
	"errors"
	"strings"
	"github.com/Xiof22/ToDoList/internal/repository"
)

type ToDoService struct {
	repo *repository.ToDoRepository
}

func NewToDoService(repo *repository.ToDoRepository) *ToDoService {
	return &ToDoService{ repo : repo }
}

func (svc *ToDoService) CreateTask(title, description string) error {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	if title == "" {
		return errors.New("Title is empty!")
	}

	svc.repo.CreateTask(title, description)
	return nil
}
