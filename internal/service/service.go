package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (models.Task, error) {
	task := models.NewTask(req.Title, req.Description)

	return svc.repo.CreateTask(ctx, task)
}

func (svc *Service) GetTasks(ctx context.Context) ([]models.Task, error) {
	return svc.repo.GetTasks(ctx)
}

func (svc *Service) GetTask(ctx context.Context, taskID models.TaskID) (models.Task, error) {
	return svc.repo.GetTask(ctx, taskID)
}

func (svc *Service) EditTask(ctx context.Context, taskID models.TaskID, req dto.EditTaskRequest) (models.Task, error) {
	task, err := svc.repo.GetTask(ctx, taskID)
	if err != nil {
		return models.Task{}, errorsx.ErrTaskNotFound
	}

	task.Title = req.Title
	task.Description = req.Description

	return svc.repo.EditTask(ctx, taskID, task)
}
