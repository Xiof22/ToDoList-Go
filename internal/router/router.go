package router

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/go-chi/chi"
)

func New(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/tasks", h.CreateTaskHandler)
	r.Get("/tasks", h.GetTasksHandler)

	return r
}
