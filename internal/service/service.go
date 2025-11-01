package service

import "github.com/Xiof22/ToDoList/internal/repository"

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}
