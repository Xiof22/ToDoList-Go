package service

import (
	"errors"
	"strings"
	"github.com/Xiof22/ToDoList/internal/repository"
)

var (
	ErrEmptyTitle = errors.New("Title is empty")
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

	if isEmptyTitle(title) {
		return ErrEmptyTitle
	}

	svc.repo.Create(title, description)
	return nil
}

func isEmptyTitle(title string) bool {
	return title == ""
}
