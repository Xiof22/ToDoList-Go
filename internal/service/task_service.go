package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateTask(ctx context.Context, req dto.CreateTaskRequest) (models.Task, error) {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return models.Task{}, ErrListNotFound
	}

	return svc.repo.CreateTask(ctx, req), nil
}

func (svc *Service) GetTasks(ctx context.Context, req dto.ListIdentifier) ([]models.Task, error) {
	if _, found := svc.repo.GetList(ctx, req); !found {
		return nil, ErrListNotFound
	}

	return svc.repo.GetTasks(ctx, req), nil
}

func (svc *Service) GetTask(ctx context.Context, req dto.TaskIdentifier) (*models.Task, bool, error) {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return nil, false, ErrListNotFound
	}

	task, found := svc.repo.GetTask(ctx, req)
	return task, found, nil
}

func (svc *Service) EditTask(ctx context.Context, req dto.EditTaskRequest) (models.Task, error) {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return models.Task{}, ErrListNotFound
	}

	if task, found := svc.repo.GetTask(ctx, dto.TaskIdentifier{
		ListID: req.ListID,
		TaskID: req.TaskID,
	}); !found {
		return models.Task{}, ErrTaskNotFound
	} else if req.Deadline.Value.Before(task.CreatedAt) && !req.Deadline.Value.IsZero() {
		return models.Task{}, ErrDeadlineBeforeCreation
	}

	return svc.repo.EditTask(ctx, req), nil
}

func (svc *Service) CompleteTask(ctx context.Context, req dto.TaskIdentifier) error {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return ErrListNotFound
	}

	if task, found := svc.repo.GetTask(ctx, req); !found {
		return ErrTaskNotFound
	} else if task.IsCompleted {
		return ErrAlreadyCompleted
	}

	svc.repo.CompleteTask(ctx, req)
	return nil
}

func (svc *Service) UncompleteTask(ctx context.Context, req dto.TaskIdentifier) error {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return ErrListNotFound
	}

	if task, found := svc.repo.GetTask(ctx, req); !found {
		return ErrTaskNotFound
	} else if !task.IsCompleted {
		return ErrAlreadyUncompleted
	}

	svc.repo.UncompleteTask(ctx, req)
	return nil
}

func (svc *Service) DeleteTask(ctx context.Context, req dto.TaskIdentifier) error {
	if _, found := svc.repo.GetList(ctx, dto.ListIdentifier{
		ID: req.ListID,
	}); !found {
		return ErrListNotFound
	}

	if _, found := svc.repo.GetTask(ctx, req); !found {
		return ErrTaskNotFound
	}

	svc.repo.DeleteTask(ctx, req)
	return nil
}
