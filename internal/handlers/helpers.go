package handlers

import (
	"fmt"
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
