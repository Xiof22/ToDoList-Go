package router

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/Xiof22/ToDoList/internal/middleware"
	"github.com/go-chi/chi"
)

func New(h *handlers.Handlers, mw *middleware.Middleware) *chi.Mux {
	r := chi.NewRouter()

	r.With(mw.AuthRequiredMiddleware).Post("/lists", h.CreateListHandler)
	r.With(mw.AuthRequiredMiddleware).Get("/lists", h.GetListsHandler)
	r.With(mw.AuthRequiredMiddleware).Get("/lists/{list_id}", h.GetListHandler)
	r.With(mw.AuthRequiredMiddleware).Patch("/lists/{list_id}", h.EditListHandler)
	r.With(mw.AuthRequiredMiddleware).Delete("/lists/{list_id}", h.DeleteListHandler)

	r.With(mw.AuthRequiredMiddleware).Post("/lists/{list_id}/tasks", h.CreateTaskHandler)
	r.With(mw.AuthRequiredMiddleware).Get("/lists/{list_id}/tasks", h.GetTasksHandler)
	r.With(mw.AuthRequiredMiddleware).Get("/lists/{list_id}/tasks/{task_id}", h.GetTaskHandler)
	r.With(mw.AuthRequiredMiddleware).Patch("/lists/{list_id}/tasks/{task_id}", h.EditTaskHandler)
	r.With(mw.AuthRequiredMiddleware).Patch("/lists/{list_id}/tasks/{task_id}/complete", h.CompleteTaskHandler)
	r.With(mw.AuthRequiredMiddleware).Patch("/lists/{list_id}/tasks/{task_id}/uncomplete", h.UncompleteTaskHandler)
	r.With(mw.AuthRequiredMiddleware).Delete("/lists/{list_id}/tasks/{task_id}", h.DeleteTaskHandler)

	r.With(mw.UnauthRequiredMiddleware).Post("/auth/register", h.RegisterHandler)
	r.With(mw.UnauthRequiredMiddleware).Post("/auth/login", h.LoginHandler)
	r.With(mw.AuthRequiredMiddleware).Post("/auth/logout", h.LogoutHandler)
	r.With(mw.AuthRequiredMiddleware).Delete("/auth/delete", h.DeleteUserHandler)

	return r
}
