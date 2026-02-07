package handlers

import (
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
	"strings"
)

// TrimStrings removes surrounding spaces from all string fields.
func trimStrings(s any) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}

	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}

func pathID[T ~[16]byte](r *http.Request, key string) (T, error) {
	raw := chi.URLParam(r, key)
	parsed, err := uuid.Parse(raw)
	if err != nil {
		var zero T
		return zero, err
	}

	return T(parsed), nil
}

func (h *Handlers) createSession(r *http.Request, w http.ResponseWriter, info models.UserInfo) error {
	session, _ := h.cs.Get(r, h.cfg.SessionName)
	session.Values[sessionKeyUserID] = info.ID.String()
	session.Values[sessionKeyUserRole] = int(info.Role)

	return session.Save(r, w)
}

func (h *Handlers) getUserInfoFromSession(r *http.Request) (models.UserInfo, error) {
	session, _ := h.cs.Get(r, h.cfg.SessionName)

	rawID, ok := session.Values[sessionKeyUserID].(string)
	if !ok {
		return models.UserInfo{}, errorsx.ErrInvalidSession
	}

	parsedID, err := uuid.Parse(rawID)
	if err != nil {
		return models.UserInfo{}, errorsx.ErrInvalidUserID
	}

	role, ok := session.Values[sessionKeyUserRole].(int)
	if !ok {
		return models.UserInfo{}, errorsx.ErrInvalidSession
	}

	info := models.UserInfo{
		ID:   models.UserID(parsedID),
		Role: models.Role(role),
	}

	return info, nil
}

func (h *Handlers) deleteSession(r *http.Request, w http.ResponseWriter) error {
	session, _ := h.cs.Get(r, h.cfg.SessionName)
	session.Options.MaxAge = -1

	return sessions.Save(r, w)
}
