package main

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/service"
)

func main() {
	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
}
