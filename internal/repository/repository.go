package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateTask(ctx context.Context, req dto.CreateTaskRequest) models.Task
}
