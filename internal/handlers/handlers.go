package handlers

import (
	"github.com/Xiof22/ToDoList/config"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/sessions"
)

const (
	sessionKeyUserID   = "user_id"
	sessionKeyUserRole = "user_role"
)

type Handlers struct {
	svc *service.Service
	cs  *sessions.CookieStore
	cfg *config.Config
}

func New(svc *service.Service, cs *sessions.CookieStore, cfg *config.Config) *Handlers {
	return &Handlers{
		svc: svc,
		cs:  cs,
		cfg: cfg,
	}
}
