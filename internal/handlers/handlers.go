package handlers

import "github.com/Xiof22/ToDoList/internal/service"

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}
