package main

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/internal/repository"
)

func main() {
	repo := repository.NewToDoRepository()
	svc := service.NewToDoService(repo)
	h := handler.NewToDoHandler(svc)
}
