package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := mux.NewRouter()

	r.HandleFunc("/tasks", h.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks", h.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", h.EditTaskHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}/complete", h.CompleteTaskHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}/uncomplete", h.UncompleteTaskHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", h.DeleteTaskHandler).Methods("DELETE")

	fmt.Println("Serving start...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
