package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

func (svc *Service) CreateTask(ctx context.Context, info models.UserInfo, listID int, req dto.CreateTaskRequest) (models.Task, error) {
	if listID <= 0 {
		return models.Task{}, errorsx.ErrInvalidListID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
	} else if list.OwnerID != info.ID {
		return models.Task{}, errorsx.ErrForbidden
	}

	task := models.NewTask(req.Title, req.Description, req.Deadline.Value)

	return svc.repo.CreateTask(ctx, listID, task), nil
}

func (svc *Service) GetTasks(ctx context.Context, info models.UserInfo, listID int) ([]models.Task, error) {
	if listID <= 0 {
		return nil, errorsx.ErrInvalidListID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return nil, err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return nil, errorsx.ErrForbidden
	}

	return svc.repo.GetTasks(ctx, listID), nil
}

func (svc *Service) GetTask(ctx context.Context, info models.UserInfo, listID, taskID int) (models.Task, error) {
	if listID <= 0 {
		return models.Task{}, errorsx.ErrInvalidListID
	}

	if taskID <= 0 {
		return models.Task{}, errorsx.ErrInvalidTaskID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.Task{}, errorsx.ErrForbidden
	}

	return svc.repo.GetTask(ctx, listID, taskID)
}

func (svc *Service) EditTask(ctx context.Context, info models.UserInfo, listID, taskID int, req dto.EditTaskRequest) (models.Task, error) {
	if listID <= 0 {
		return models.Task{}, errorsx.ErrInvalidListID
	}

	if taskID <= 0 {
		return models.Task{}, errorsx.ErrInvalidTaskID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return models.Task{}, errorsx.ErrForbidden
	}

	task, err := svc.repo.GetTask(ctx, listID, taskID)
	if err != nil {
		return task, err
	}

	if req.Deadline.Value.Before(task.CreatedAt) && !req.Deadline.Value.IsZero() {
		return models.Task{}, errorsx.ErrDeadlineBeforeCreation
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Deadline = req.Deadline.Value
	task.UpdatedAt = time.Now()

	return task, svc.repo.EditTask(ctx, listID, taskID, task)
}

func (svc *Service) CompleteTask(ctx context.Context, info models.UserInfo, listID, taskID int) error {
	if listID <= 0 {
		return errorsx.ErrInvalidListID
	}

	if taskID <= 0 {
		return errorsx.ErrInvalidTaskID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	} else if list.OwnerID != info.ID {
		return errorsx.ErrForbidden
	}

	task, err := svc.repo.GetTask(ctx, listID, taskID)
	if err != nil {
		return err
	}

	if task.IsCompleted {
		return errorsx.ErrAlreadyCompleted
	}

	task.IsCompleted = true

	return svc.repo.EditTask(ctx, listID, taskID, task)
}

func (svc *Service) UncompleteTask(ctx context.Context, info models.UserInfo, listID, taskID int) error {
	if listID <= 0 {
		return errorsx.ErrInvalidListID
	}

	if taskID <= 0 {
		return errorsx.ErrInvalidTaskID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	} else if list.OwnerID != info.ID {
		return errorsx.ErrForbidden
	}

	task, err := svc.repo.GetTask(ctx, listID, taskID)
	if err != nil {
		return err
	}

	if !task.IsCompleted {
		return errorsx.ErrAlreadyUncompleted
	}

	task.IsCompleted = false

	return svc.repo.EditTask(ctx, listID, taskID, task)
}

func (svc *Service) DeleteTask(ctx context.Context, info models.UserInfo, listID, taskID int) error {
	if listID <= 0 {
		return errorsx.ErrInvalidListID
	}

	if taskID <= 0 {
		return errorsx.ErrInvalidTaskID
	}

	if list, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	} else if list.OwnerID != info.ID && info.Role != models.Admin {
		return errorsx.ErrForbidden
	}

	if _, err := svc.repo.GetTask(ctx, listID, taskID); err != nil {
		return err
	}

	svc.repo.DeleteTask(ctx, listID, taskID)
	return nil
}
