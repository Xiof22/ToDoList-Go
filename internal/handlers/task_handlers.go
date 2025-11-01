package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

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

	task, err := h.svc.CreateTask(r.Context(), listID, req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	tasks, err := h.svc.GetTasks(r.Context(), listID)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.TasksResponse{
		Count: len(tasks),
		Tasks: dto.ToTaskDTOs(tasks),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	taskID, err := pathID[models.TaskID](r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	task, err := h.svc.GetTask(r.Context(), listID, taskID)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	taskID, err := pathID[models.TaskID](r, pathKeyTaskID)
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

	task, err := h.svc.EditTask(r.Context(), listID, taskID, req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	taskID, err := pathID[models.TaskID](r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	if err := h.svc.CompleteTask(r.Context(), listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) UncompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	taskID, err := pathID[models.TaskID](r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	if err := h.svc.UncompleteTask(r.Context(), listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	taskID, err := pathID[models.TaskID](r, pathKeyTaskID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidTaskID)
		return
	}

	if err := h.svc.DeleteTask(r.Context(), listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
