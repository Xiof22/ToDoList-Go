package router

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/go-chi/chi"
)

func New(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/lists", h.CreateListHandler)
	r.Get("/lists", h.GetListsHandler)
	r.Get("/lists/{list_id}", h.GetListHandler)
	r.Patch("/lists/{list_id}", h.EditListHandler)
	r.Delete("/lists/{list_id}", h.DeleteListHandler)

	r.Post("/lists/{list_id}/tasks", h.CreateTaskHandler)
	r.Get("/lists/{list_id}/tasks", h.GetTasksHandler)
	r.Get("/lists/{list_id}/tasks/{task_id}", h.GetTaskHandler)
	r.Patch("/lists/{list_id}/tasks/{task_id}", h.EditTaskHandler)
	r.Patch("/lists/{list_id}/tasks/{task_id}/complete", h.CompleteTaskHandler)
	r.Patch("/lists/{list_id}/tasks/{task_id}/uncomplete", h.UncompleteTaskHandler)
	r.Delete("/lists/{list_id}/tasks/{task_id}", h.DeleteTaskHandler)

	return r
}
