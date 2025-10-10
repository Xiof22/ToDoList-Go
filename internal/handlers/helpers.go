package handlers

import (
	"errors"
	"fmt"
	"github.com/Xiof22/ToDoList/internal/service"
	"github.com/go-chi/chi"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// TrimStrings removes surrounding spaces from all string fields.
func trimStrings(s any) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}

	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}

func getURLIntParam(r *http.Request, key string) (int, error) {
	paramStr := chi.URLParam(r, key)
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse '%s'", key)
	}

	return id, nil
}

func mapTaskError(err error) int {
	switch {
	case errors.Is(err, service.ErrAlreadyCompleted),
		errors.Is(err, service.ErrAlreadyUncompleted):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrTaskNotFound):
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
