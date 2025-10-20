package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	setTimezone("Asia/Ashgabat")

	repo := repository.NewToDoRepository()
	svc := service.NewToDoService(repo)
	h := handler.NewToDoHandler(svc)
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

func setTimezone(location string) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Println("Location loading error:", err)
		fmt.Println("Leaving default timezone (UTC +0000)")
		return
	}

	time.Local = loc
}
