package handler

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
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
	writeResponse(w, "Created succefully")
}

func (h *ToDoHandler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := h.svc.GetTasks()

	if len(tasks) == 0 {
		writeResponse(w, "There's no tasks")
		return
	}

	for _, task := range tasks {
		writeResponse(w, task)
	}
}

func (h *ToDoHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "ID parsing error", http.StatusBadRequest)
		return
	}

	task, err := h.svc.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		writeResponse(w, "Task not found")
		return
	}

	writeResponse(w, task)
}

func (h *ToDoHandler) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "ID parsing error", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")

	err = h.svc.EditTask(id, title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResponse(w, "Edited succefully")
}

func (h *ToDoHandler) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "ID parsing error", http.StatusBadRequest)
		return
	}

	err = h.svc.CompleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeResponse(w, "Completed task succefully")
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("content-type", "text/plain")
	switch v := data.(type) {

		case string :
			fmt.Fprintf(w, v)

		case fmt.Stringer :
			fmt.Fprintf(w, v.String())

	}
}

func parseID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	return strconv.Atoi(idStr)
}
