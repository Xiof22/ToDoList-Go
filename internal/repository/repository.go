package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
	"time"
)

type Repository interface {
	Create(title, description string, deadline time.Time) error
	GetAll() ([]models.Task, error)
	Get(id int) (*models.Task, error)
	Edit(id int, title, description string, deadline time.Time) error
	Complete(id int) error
	Uncomplete(id int) error
	Delete(id int) error
}
