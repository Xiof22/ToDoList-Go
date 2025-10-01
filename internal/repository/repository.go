package repository

import (
	"github.com/Xiof22/ToDoList/internal/models"
)

type ToDoRepository struct {
	Tasks []models.Task
	nextID int
}

func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{ nextID : 1 }
}
