package main

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/repository/memory"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/internal/validator"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	setTimezone("Asia/Ashgabat")

	m := memory.New()
	svc := service.New(m)
	h := handlers.New(svc)
	r := mux.NewRouter()
	validator.Init()

	// List Handlers
	r.HandleFunc("/lists", h.CreateListHandler).Methods("POST")
	r.HandleFunc("/lists", h.GetListsHandler).Methods("GET")
	r.HandleFunc("/lists/{list_id}", h.GetListHandler).Methods("GET")
	r.HandleFunc("/lists/{list_id}", h.EditListHandler).Methods("PATCH")
	r.HandleFunc("/lists/{list_id}", h.DeleteListHandler).Methods("DELETE")

	// Task Handlers
	r.HandleFunc("/lists/{list_id}/tasks", h.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/lists/{list_id}/tasks", h.GetTasksHandler).Methods("GET")
	r.HandleFunc("/lists/{list_id}/tasks/{task_id}", h.GetTaskHandler).Methods("GET")
	r.HandleFunc("/lists/{list_id}/tasks/{task_id}", h.EditTaskHandler).Methods("PATCH")
	r.HandleFunc("/lists/{list_id}/tasks/{task_id}/complete", h.CompleteTaskHandler).Methods("PATCH")
	r.HandleFunc("/lists/{list_id}/tasks/{task_id}/uncomplete", h.UncompleteTaskHandler).Methods("PATCH")
	r.HandleFunc("/lists/{list_id}/tasks/{task_id}", h.DeleteTaskHandler).Methods("DELETE")

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
