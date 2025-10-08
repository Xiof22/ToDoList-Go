package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

type Repository interface {
	Create(title, description string)
	GetAll() ([]models.Task, error)
	Get(id int) (*models.Task, error)
}
