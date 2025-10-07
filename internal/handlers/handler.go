package handler

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/service"
	"net/http"
	"strings"
)

const (
	errInvalidTitle = "Invalid title"
)

type ToDoHandler struct {
	svc *service.ToDoService
}

func NewToDoHandler(svc *service.ToDoService) *ToDoHandler {
	return &ToDoHandler{svc: svc}
}

func (h *ToDoHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimSpace(r.FormValue("title"))
	description := strings.TrimSpace(r.FormValue("description"))

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	h.svc.CreateTask(title, description)

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, "Created succefully")
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "text/plain")
	switch v := data.(type) {

	case string:
		fmt.Fprintf(w, v)

	}
}

func isEmpty(str string) bool {
	return str == ""
}
