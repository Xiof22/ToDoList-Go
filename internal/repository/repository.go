package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateList(ctx context.Context, list models.List) (models.List, error)
	GetLists(ctx context.Context) ([]models.List, error)
	GetListsByUserID(ctx context.Context, userID models.UserID) ([]models.List, error)
	GetList(ctx context.Context, listID models.ListID) (models.List, error)
	EditList(ctx context.Context, listID models.ListID, list models.List) (models.List, error)
	DeleteList(ctx context.Context, listID models.ListID) error

	CreateTask(ctx context.Context, listID models.ListID, task models.Task) (models.Task, error)
	GetTasks(ctx context.Context, listID models.ListID) ([]models.Task, error)
	GetTask(ctx context.Context, listID models.ListID, taskID models.TaskID) (models.Task, error)
	EditTask(ctx context.Context, listID models.ListID, taskID models.TaskID, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, listID models.ListID, taskID models.TaskID) error

	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, userID models.UserID) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	DeleteUser(ctx context.Context, userID models.UserID) error
}
