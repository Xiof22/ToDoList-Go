package handlers

import (
	"fmt"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

const (
	errInvalidTitle = "Invalid title"
	errInvalidID    = "Invalid ID"
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

	err := h.svc.CreateTask(title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created succefully")
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.svc.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		fmt.Fprint(w, "There's no tasks")
		return
	}

	for _, task := range tasks {
		fmt.Fprint(w, task)
	}
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil || !isPositive(id) {
		http.Error(w, errInvalidID, http.StatusBadRequest)
		return
	}

	task, err := h.svc.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Task not found")
		return
	}

	fmt.Fprint(w, task)
}

func (h *Handlers) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil || !isPositive(id) {
		http.Error(w, errInvalidID, http.StatusBadRequest)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	description := strings.TrimSpace(r.FormValue("description"))

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	err = h.svc.EditTask(id, title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Edited succefully")
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil || !isPositive(id) {
		http.Error(w, errInvalidID, http.StatusBadRequest)
		return
	}

	err = h.svc.CompleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Completed task succefully")
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
