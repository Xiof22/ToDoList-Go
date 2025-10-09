package router

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/go-chi/chi"
)

func New(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/tasks", h.CreateTaskHandler)
	r.Get("/tasks", h.GetTasksHandler)
	r.Get("/tasks/{task_id}", h.GetTaskHandler)
	r.Patch("/tasks/{task_id}", h.EditTaskHandler)

	return r
}
