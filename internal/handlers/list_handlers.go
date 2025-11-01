package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) CreateListHandler(w http.ResponseWriter, r *http.Request) {
	title := getFormValueWithTrim(r, "title")
	description := getFormValueWithTrim(r, "description")

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	err := h.svc.CreateList(title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created list succefully")
}

func (h *Handlers) GetListsHandler(w http.ResponseWriter, r *http.Request) {
	lists, err := h.svc.GetLists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(lists) == 0 {
		fmt.Fprint(w, "There`s no lists")
		return
	}

	var response string
	for _, list := range lists {
		response += list.String()
	}

	fmt.Fprint(w, response)
}

func (h *Handlers) GetListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	list, err := h.svc.GetList(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, list)
}

func (h *Handlers) EditListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	title := getFormValueWithTrim(r, "title")
	description := getFormValueWithTrim(r, "description")

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	err = h.svc.EditList(listID, title, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Edited list succefully")
}

func (h *Handlers) DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	err = h.svc.DeleteList(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Deleted list succefully")
}
