package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
}
