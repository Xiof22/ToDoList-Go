package middleware

import (
	"github.com/Xiof22/ToDoList/config"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/gorilla/sessions"
	"net/http"
)

type Middleware struct {
	cs  *sessions.CookieStore
	cfg *config.Config
}

func New(cs *sessions.CookieStore, cfg *config.Config) *Middleware {
	return &Middleware{
		cs:  cs,
		cfg: cfg,
	}
}

func (mw *Middleware) AuthRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := mw.cs.Get(r, mw.cfg.SessionName)
		if session.IsNew {
			responses.WriteError(w, http.StatusUnauthorized, errorsx.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mw *Middleware) UnauthRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := mw.cs.Get(r, mw.cfg.SessionName)
		if !session.IsNew {
			responses.WriteError(w, http.StatusBadRequest, errorsx.ErrAlreadyAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
