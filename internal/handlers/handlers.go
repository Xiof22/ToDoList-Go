package handlers

import (
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

const (
	errInvalidTitle    = "Invalid title"
	errInvalidListID   = "Invalid list ID"
	errInvalidTaskID   = "Invalid task ID"
	errInvalidDeadline = "Unexpected deadline format"
	timeLayout         = "2006-01-02T15:04"
)

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}

func getFormValueWithTrim(r *http.Request, key string) string {
	return strings.TrimSpace(r.FormValue(key))
}

func isEmpty(str string) bool {
	return str == ""
}

func isPositive(n int) bool {
	return n > 0
}

func getURLIntParam(r *http.Request, key string) (int, error) {
	vars := mux.Vars(r)
	paramStr := vars[key]
	return strconv.Atoi(paramStr)
}
