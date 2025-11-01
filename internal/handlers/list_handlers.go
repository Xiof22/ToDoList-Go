package handlers

import (
	"encoding/json"
	"github.com/Xiof22/ToDoList/internal/dto"
	"github.com/Xiof22/ToDoList/internal/validator"
	"net/http"
)

func (h *Handlers) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	list := h.svc.CreateList(r.Context(), req)
	resp := dto.ListResponse{
		List: dto.ToListDTO(&list),
	}

	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handlers) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	lists := h.svc.GetLists(r.Context())
	resp := dto.ListsResponse{
		Count: len(lists),
		Lists: dto.ToListDTOs(lists),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) GetListHandler(w http.ResponseWriter, r *http.Request) {
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

	list, found := h.svc.GetList(r.Context(), req)

	status := http.StatusNotFound
	if found {
		status = http.StatusOK
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(list),
	}

	writeJSON(w, status, resp)
}

func (h *Handlers) EditListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	req := dto.EditListRequest{ListID: listID}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	trimStrings(&req)
	if err := validator.Validate.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	list, err := h.svc.EditList(r.Context(), req)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	resp := dto.ListResponse{
		List: dto.ToListDTO(&list),
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *Handlers) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := h.svc.DeleteList(r.Context(), req); err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
