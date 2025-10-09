package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	Create(title, description string) error
	GetAll() ([]models.Task, error)
	Get(id int) (*models.Task, error)
	Edit(id int, title, description string) error
}
