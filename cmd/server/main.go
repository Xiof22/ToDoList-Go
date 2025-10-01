package main

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository"
	"github.com/Xiof22/ToDoList/internal/service"
)

func main() {
	repo := repository.NewToDoRepository()
	svc := service.NewToDoService(repo)
	h := handler.NewToDoHandler(svc)
}
