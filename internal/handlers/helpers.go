package handlers

import (
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
	"strconv"
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

func getURLIntParam(r *http.Request, key string) (int, error) {
	paramStr := chi.URLParam(r, key)
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, errorsx.ErrParseURL(key)
	}

	return id, nil
}

func (h *Handlers) createSession(r *http.Request, w http.ResponseWriter, info models.UserInfo) error {
	session, _ := h.cs.Get(r, h.cfg.SessionName)
	session.Values[sessionKeyUserID] = info.ID
	session.Values[sessionKeyUserRole] = int(info.Role)

	return session.Save(r, w)
}

func (h *Handlers) getUserInfoFromSession(r *http.Request) (models.UserInfo, error) {
	session, _ := h.cs.Get(r, h.cfg.SessionName)

	id, ok := session.Values[sessionKeyUserID].(int)
	if !ok {
		return models.UserInfo{}, errorsx.ErrInvalidSession
	}

	role, ok := session.Values[sessionKeyUserRole].(int)
	if !ok {
		return models.UserInfo{}, errorsx.ErrInvalidSession
	}

	info := models.UserInfo{
		ID:   id,
		Role: models.Role(role),
	}

	return info, nil
}

func (h *Handlers) deleteSession(r *http.Request, w http.ResponseWriter) error {
	session, _ := h.cs.Get(r, h.cfg.SessionName)
	session.Options.MaxAge = -1

	return sessions.Save(r, w)
}
