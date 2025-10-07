package handlers

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/service"
	"net/http"
	"strings"
)

const (
	errInvalidTitle = "Invalid title"
)

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := getFormValueWithTrim(r, "title")
	description := getFormValueWithTrim(r, "description")

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	h.svc.CreateTask(title, description)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created succefully")
}

func getFormValueWithTrim(r *http.Request, key string) string {
	return strings.TrimSpace(r.FormValue(key))
}

func isEmpty(str string) bool {
	return str == ""
}
