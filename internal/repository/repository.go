package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/models"
)

//go:generate mockery --name Repository
type Repository interface {
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
}
