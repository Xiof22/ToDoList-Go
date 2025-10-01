package service

import (
	"github.com/Xiof22/ToDoList/internal/repository"
)

type ToDoService struct {
	repo *repository.ToDoRepository
}

func NewToDoService(repo *repository.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}
