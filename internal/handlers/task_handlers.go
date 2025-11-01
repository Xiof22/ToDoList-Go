package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	title := getFormValueWithTrim(r, "title")
	description := getFormValueWithTrim(r, "description")
	deadlineStr := getFormValueWithTrim(r, "deadline")

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	deadline, err := time.Parse(timeLayout, deadlineStr)
	if err != nil && !isEmpty(deadlineStr) {
		http.Error(w, errInvalidDeadline, http.StatusBadRequest)
		return
	}

	err = h.svc.CreateTask(listID, title, description, deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created task succefully")
}

func (h *Handlers) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	tasks, err := h.svc.GetTasks(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		fmt.Fprint(w, "There's no tasks")
		return
	}

	for _, task := range tasks {
		fmt.Fprint(w, task)
	}
}

func (h *Handlers) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil || !isPositive(taskID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	task, err := h.svc.GetTask(listID, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Task not found")
		return
	}

	fmt.Fprint(w, task)
}

func (h *Handlers) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil || !isPositive(taskID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	title := getFormValueWithTrim(r, "title")
	description := getFormValueWithTrim(r, "description")
	deadlineStr := getFormValueWithTrim(r, "deadline")

	if isEmpty(title) {
		http.Error(w, errInvalidTitle, http.StatusBadRequest)
		return
	}

	deadline, err := time.Parse(timeLayout, deadlineStr)
	if err != nil && !isEmpty(deadlineStr) {
		http.Error(w, errInvalidDeadline, http.StatusBadRequest)
		return
	}

	err = h.svc.EditTask(listID, taskID, title, description, deadline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Edited task succefully")
}

func (h *Handlers) CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil || !isPositive(taskID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	err = h.svc.CompleteTask(listID, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Completed task succefully")
}

func (h *Handlers) UncompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil || !isPositive(taskID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	err = h.svc.UncompleteTask(listID, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Uncompleted task succefully")
}

func (h *Handlers) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	listID, err := getURLIntParam(r, "list_id")
	if err != nil || !isPositive(listID) {
		http.Error(w, errInvalidListID, http.StatusBadRequest)
		return
	}

	taskID, err := getURLIntParam(r, "task_id")
	if err != nil || !isPositive(taskID) {
		http.Error(w, errInvalidTaskID, http.StatusBadRequest)
		return
	}

	err = h.svc.DeleteTask(listID, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Deleted succefully")
}
