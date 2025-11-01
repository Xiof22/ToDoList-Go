package handlers

import "github.com/Xiof22/ToDoList/internal/service"

const (
	pathKeyListID = "list_id"
	pathKeyTaskID = "task_id"
)

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}
