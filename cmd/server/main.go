package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	repo := repository.NewToDoRepository()
	svc := service.NewToDoService(repo)
	h := handler.NewToDoHandler(svc)
	r := mux.NewRouter()

	r.HandleFunc("/tasks", h.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", h.GetTasksHandler).Methods("GET")

	fmt.Println("Serving start...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
