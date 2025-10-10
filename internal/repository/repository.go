package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task
	GetTasks(ctx context.Context) []models.Task
	GetTask(ctx context.Context, req dto.TaskIdentifier) (*models.Task, bool)
	EditTask(ctx context.Context, dto dto.EditTaskRequest) models.Task
	CompleteTask(ctx context.Context, req dto.TaskIdentifier)
	UncompleteTask(ctx context.Context, req dto.TaskIdentifier)
}
