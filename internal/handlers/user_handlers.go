package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/errorsx"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.svc.Register(r.Context(), req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	if err := h.createSession(r, w, user.Info()); err != nil {
		responses.WriteError(w, http.StatusInternalServerError, errorsx.ErrSaveSession)
		return
	}

	resp := dto.UserResponse{
		User: dto.ToUserDTO(user),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.svc.Login(r.Context(), req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	if err := h.createSession(r, w, user.Info()); err != nil {
		fmt.Println(err)
		responses.WriteError(w, http.StatusInternalServerError, errorsx.ErrSaveSession)
		return
	}

	resp := dto.UserResponse{
		User: dto.ToUserDTO(user),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.deleteSession(r, w); err != nil {
		responses.WriteError(w, http.StatusInternalServerError, errorsx.ErrSaveSession)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.DeleteUser(r.Context(), info); err != nil {
		responses.WriteError(w, http.StatusNotFound, err)
		return
	}

	if err := h.deleteSession(r, w); err != nil {
		responses.WriteError(w, http.StatusInternalServerError, errorsx.ErrSaveSession)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
