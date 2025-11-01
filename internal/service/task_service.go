package service

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
)

func (svc *Service) CreateTask(ctx context.Context, listID models.ListID, req dto.CreateTaskRequest) (models.Task, error) {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
	}

	task := models.NewTask(req.Title, req.Description, req.Deadline.Value)

	return svc.repo.CreateTask(ctx, listID, task)
}

func (svc *Service) GetTasks(ctx context.Context, listID models.ListID) ([]models.Task, error) {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return nil, err
	}

	return svc.repo.GetTasks(ctx, listID)
}

func (svc *Service) GetTask(ctx context.Context, listID models.ListID, taskID models.TaskID) (models.Task, error) {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
	}

	return svc.repo.GetTask(ctx, listID, taskID)
}

func (svc *Service) EditTask(ctx context.Context, listID models.ListID, taskID models.TaskID, req dto.EditTaskRequest) (models.Task, error) {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return models.Task{}, err
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

	return svc.repo.EditTask(ctx, listID, taskID, task)
}

func (svc *Service) CompleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	}

	task, err := svc.repo.GetTask(ctx, listID, taskID)
	if err != nil {
		return err
	}

	if task.IsCompleted {
		return errorsx.ErrAlreadyCompleted
	}

	task.IsCompleted = true

	_, err = svc.repo.EditTask(ctx, listID, taskID, task)
	return err
}

func (svc *Service) UncompleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	}

	task, err := svc.repo.GetTask(ctx, listID, taskID)
	if err != nil {
		return err
	}

	if !task.IsCompleted {
		return errorsx.ErrAlreadyUncompleted
	}

	task.IsCompleted = false

	_, err = svc.repo.EditTask(ctx, listID, taskID, task)
	return err
}

func (svc *Service) DeleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error {
	if _, err := svc.repo.GetList(ctx, listID); err != nil {
		return err
	}

	if _, err := svc.repo.GetTask(ctx, listID, taskID); err != nil {
		return err
	}

	return svc.repo.DeleteTask(ctx, listID, taskID)
}
