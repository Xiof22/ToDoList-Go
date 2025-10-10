package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task := h.svc.CreateTask(r.Context(), req)
	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(&task),
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := h.svc.GetTasks(r.Context())

	writeJSON(w, http.StatusOK, dto.TasksResponse{
		Count: len(tasks),
		Tasks: dto.ToTaskDTOs(tasks),
	})
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{ID: id}
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, found := h.svc.GetTask(r.Context(), req)
	status := http.StatusNotFound
	if found {
		status = http.StatusOK
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	writeJSON(w, status, resp)
}

func (h *Handlers) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.EditTaskRequest{ID: id}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.svc.EditTask(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(&task),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{ID: id}
	if err := validator.Validate.Struct(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.CompleteTask(r.Context(), req); err != nil {
		writeError(w, mapTaskError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) UncompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getURLIntParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{ID: id}
	if err := validator.Validate.Struct(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.UncompleteTask(r.Context(), req); err != nil {
		writeError(w, mapTaskError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
