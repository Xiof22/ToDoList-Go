package handler

import (
	"github.com/Xiof22/ToDoList/internal/service"
)

type ToDoHandler struct {
	svc *service.ToDoService
}

func NewToDoHandler(svc *service.ToDoService) *ToDoHandler {
	return &ToDoHandler{ svc : svc }
}
