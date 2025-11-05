package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTask(ctx context.Context, taskID models.TaskID) (models.Task, error)
	EditTask(ctx context.Context, taskID models.TaskID, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, taskID models.TaskID) error
}
