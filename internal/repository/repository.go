package repository

import (
	"context"
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	CreateList(ctx context.Context, list models.List) models.List
	GetLists(ctx context.Context) []models.List
	GetListsByUserID(ctx context.Context, userID int) []models.List
	GetList(ctx context.Context, listID int) (models.List, error)
	EditList(ctx context.Context, listID int, list models.List) error
	DeleteList(ctx context.Context, listID int)

	CreateTask(ctx context.Context, listID int, task models.Task) models.Task
	GetTasks(ctx context.Context, listID int) []models.Task
	GetTask(ctx context.Context, listID, taskID int) (models.Task, error)
	EditTask(ctx context.Context, listID, taskID int, task models.Task) error
	//	CompleteTask(ctx context.Context, listID, taskID int)
	//	UncompleteTask(ctx context.Context, listID, taskID int)
	DeleteTask(ctx context.Context, listID, taskID int)

	CreateUser(ctx context.Context, user models.User) models.User
	GetUserByID(ctx context.Context, userID int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	DeleteUser(ctx context.Context, userID int) error
}
