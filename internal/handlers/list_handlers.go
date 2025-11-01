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

func (h *Handlers) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	list, err := h.svc.CreateList(r.Context(), req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(list),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := h.svc.GetLists(r.Context())
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.ListsResponse{
		Count: len(lists),
		Lists: dto.ToListDTOs(lists),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	list, err := h.svc.GetList(r.Context(), listID)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(list),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) EditListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	var req dto.EditListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	list, err := h.svc.EditList(r.Context(), listID, req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(list),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := pathID[models.ListID](r, pathKeyListID)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, errorsx.ErrInvalidListID)
		return
	}

	if err := h.svc.DeleteList(r.Context(), listID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
