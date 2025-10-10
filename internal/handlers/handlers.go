package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

const pathKeyTaskID = "task_id"

type Handlers struct {
	svc *service.Service
}

func New(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.svc.CreateTask(r.Context(), req)
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.svc.GetTasks(r.Context())
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	responses.WriteJSON(w, http.StatusOK, dto.TasksResponse{
		Count: len(tasks),
		Tasks: dto.ToTaskDTOs(tasks),
	})
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID, err := pathID(r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	task, err := h.svc.GetTask(r.Context(), taskID)
	if err != nil {
		responses.WriteError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID, err := pathID(r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	var req dto.EditTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.svc.EditTask(r.Context(), taskID, req)
	if err != nil {
		responses.WriteError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID, err := pathID(r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	if err := h.svc.CompleteTask(r.Context(), taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) UncompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID, err := pathID(r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	if err := h.svc.UncompleteTask(r.Context(), taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
