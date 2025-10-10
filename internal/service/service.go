package service

import (
	"errors"
	"strings"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
)

var (
	ErrEmptyTitle = errors.New("Title is empty")
	ErrInvalidID = errors.New("Invalid ID")
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

func (svc *ToDoService) GetTasks() []models.Task {
	return svc.repo.GetAll()
}

func (svc *ToDoService) GetTask(id int) (*models.Task, error) {
	if !isValidID(id) {
		return nil, ErrInvalidID
	}

	return svc.repo.Get(id), nil
}

func (svc *ToDoService) EditTask(id int, title, description string) error {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	if !isValidID(id) {
		return ErrInvalidID
	}

	if isEmptyTitle(title) {
		return ErrEmptyTitle
	}

	return svc.repo.Edit(id, title, description)
}

func (svc *ToDoService) CompleteTask(id int) error {
	if !isValidID(id) {
		return ErrInvalidID
	}

	return svc.repo.Complete(id)
}

func (svc *ToDoService) UncompleteTask(id int) error {
	if !isValidID(id) {
		return ErrInvalidID
	}

	return svc.repo.Uncomplete(id)
}

func isEmptyTitle(title string) bool {
	return title == ""
}

func isValidID(id int) bool {
	return id > 0
}
