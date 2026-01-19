package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
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

	task, err := h.svc.CreateTask(r.Context(), info, listID, req)
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
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	tasks, err := h.svc.GetTasks(r.Context(), info, listID)
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
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.svc.GetTask(r.Context(), info, listID, taskID)
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
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
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

	task, err := h.svc.EditTask(r.Context(), info, listID, taskID, req)
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
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.CompleteTask(r.Context(), info, listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) UncompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.UncompleteTask(r.Context(), info, listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.DeleteTask(r.Context(), info, listID, taskID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
