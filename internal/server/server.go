package server

import (
	"github.com/Xiof22/ToDoList/config"
	"github.com/go-chi/chi"
	"net/http"
)

func New(r *chi.Mux, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:         cfg.Addr,
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
