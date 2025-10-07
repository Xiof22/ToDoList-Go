package router

import (
	"github.com/Xiof22/ToDoList/internal/handlers"
	"github.com/go-chi/chi"
)

func New(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	return r
}
