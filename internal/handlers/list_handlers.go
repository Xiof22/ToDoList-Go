package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/models"
	"github.com/Xiof22/ToDoList/internal/responses"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

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

	list, err := h.svc.CreateList(r.Context(), info, req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(list, false),
	}

	responses.WriteJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	info, err := h.getUserInfoFromSession(r)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	lists := h.svc.GetLists(r.Context(), info)
	withOwnerID := info.Role == models.Admin
	resp := dto.ListsResponse{
		Count: len(lists),
		Lists: dto.ToListDTOs(lists, withOwnerID),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetListHandler(w http.ResponseWriter, r *http.Request) {
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

	list, err := h.svc.GetList(r.Context(), info, listID)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	withOwnerID := info.ID != list.OwnerID
	resp := dto.ListResponse{
		List: dto.ToListDTO(list, withOwnerID),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) EditListHandler(w http.ResponseWriter, r *http.Request) {
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

	list, err := h.svc.EditList(r.Context(), info, listID, req)
	if err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	withOwnerID := list.OwnerID != info.ID
	resp := dto.ListResponse{
		List: dto.ToListDTO(list, withOwnerID),
	}

	responses.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := h.svc.DeleteList(r.Context(), info, listID); err != nil {
		responses.WriteError(w, responses.MapError(err), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
