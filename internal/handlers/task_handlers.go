package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.CreateTaskRequest{ListID: listID}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.svc.CreateTask(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(&task),
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.ListIdentifier{ID: listID}
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	tasks, err := h.svc.GetTasks(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.TasksResponse{
		Count: len(tasks),
		Tasks: dto.ToTaskDTOs(tasks),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{
		ListID: listID,
		TaskID: taskID,
	}

	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	task, found, err := h.svc.GetTask(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

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
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.EditTaskRequest{
		ListID: listID,
		TaskID: taskID,
	}

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
		writeError(w, mapTaskError(err), err)
		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(&task),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{
		ListID: listID,
		TaskID: taskID,
	}

	if err := validator.Validate.Struct(req); err != nil {
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
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{
		ListID: listID,
		TaskID: taskID,
	}

	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.UncompleteTask(r.Context(), req); err != nil {
		writeError(w, mapTaskError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.TaskIdentifier{
		ListID: listID,
		TaskID: taskID,
	}

	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.DeleteTask(r.Context(), req); err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
