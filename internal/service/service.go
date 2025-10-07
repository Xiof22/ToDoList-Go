package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task {
	return svc.repo.CreateTask(ctx, req)
}
