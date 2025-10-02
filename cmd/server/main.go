package main

import (
	"net/http"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/internal/repository"
)

func main() {
	repo := repository.NewToDoRepository()
	svc := service.NewToDoService(repo)
	h := handler.NewToDoHandler(svc)

	http.HandleFunc("POST /tasks", h.CreateTaskHandler)
	http.ListenAndServe(":8080", nil)
}
