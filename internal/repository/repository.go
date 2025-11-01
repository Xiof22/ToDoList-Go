package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

type Repository interface {

	// List
	CreateList(title, description string) error
	GetLists() ([]models.List, error)
	GetList(listID int) (*models.List, error)
	EditList(listID int, title, description string) error
	DeleteList(listID int) error

	// Task
	CreateTask(listID int, title, description string, deadline time.Time) error
	GetTasks(listID int) ([]models.Task, error)
	GetTask(listID, taskID int) (*models.Task, error)
	EditTask(listID, taskID int, title, description string, deadline time.Time) error
	CompleteTask(listID, taskID int) error
	UncompleteTask(listID, taskID int) error
	DeleteTask(listID, taskID int) error
}
