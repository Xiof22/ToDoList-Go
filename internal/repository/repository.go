package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateList(ctx context.Context, req dto.CreateListRequest) models.List
	GetLists(ctx context.Context) []models.List
	GetList(ctx context.Context, req dto.ListIdentifier) (*models.List, bool)
	EditList(ctx context.Context, req dto.EditListRequest) models.List
	DeleteList(ctx context.Context, req dto.ListIdentifier)

	CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task
	GetTasks(ctx context.Context, req dto.ListIdentifier) []models.Task
	GetTask(ctx context.Context, req dto.TaskIdentifier) (*models.Task, bool)
	EditTask(ctx context.Context, dto dto.EditTaskRequest) models.Task
	CompleteTask(ctx context.Context, req dto.TaskIdentifier)
	UncompleteTask(ctx context.Context, req dto.TaskIdentifier)
	DeleteTask(ctx context.Context, req dto.TaskIdentifier)
}
