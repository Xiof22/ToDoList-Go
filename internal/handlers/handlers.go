package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/responses"
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
		if !errors.Is(err, context.Canceled) {
			responses.WriteError(w, http.StatusInternalServerError, err)
		}

		return
	}

	resp := dto.TaskResponse{
		Task: dto.ToTaskDTO(task),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}
