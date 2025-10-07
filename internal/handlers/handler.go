package handler

import (
	"fmt"
	"net/http"
	"github.com/Xiof22/ToDoList/internal/service"
)

type ToDoHandler struct {
	svc *service.ToDoService
}

func NewToDoHandler(svc *service.ToDoService) *ToDoHandler {
	return &ToDoHandler{ svc : svc }
}

func (h *ToDoHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	err := h.svc.CreateTask(title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeResponse(w, "Created succefully!")
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "text/plain")
	switch v := data.(type) {

		case string :
			fmt.Fprintf(w, v)

	}
}
